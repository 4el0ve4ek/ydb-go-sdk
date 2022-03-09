package conn

import (
	"context"
	"sync"
	"sync/atomic"
	"time"

	"google.golang.org/grpc"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/endpoint"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/errors"
	"github.com/ydb-platform/ydb-go-sdk/v3/trace"
)

type Pool interface {
	PoolGetter
	TakerReleaser
	Pessimizer
}

type PoolGetter interface {
	GetConn(endpoint endpoint.Endpoint) Conn
}

type TakerReleaser interface {
	Take(ctx context.Context) error
	Release(ctx context.Context) error
}

type Pessimizer interface {
	Pessimize(ctx context.Context, cc Conn, cause error)
}

type PoolConfig interface {
	ConnectionTTL() time.Duration
	GrpcDialOptions() []grpc.DialOption
}

type pool struct {
	usages int64
	config Config
	mtx    sync.RWMutex
	opts   []grpc.DialOption
	conns  map[string]Conn
	done   chan struct{}
}

func (p *pool) Pessimize(ctx context.Context, cc Conn, cause error) {
	e := cc.Endpoint().Copy()
	cc, ok := p.conns[e.Address()]
	if !ok {
		return
	}

	trace.DriverOnPessimizeNode(
		trace.ContextDriver(ctx).Compose(p.config.Trace()),
		&ctx,
		e,
		cc.GetState(),
		cause,
	)(cc.SetState(Banned))
}

func (p *pool) Take(ctx context.Context) error {
	atomic.AddInt64(&p.usages, 1)
	return nil
}

func (p *pool) Release(ctx context.Context) error {
	if atomic.AddInt64(&p.usages, -1) > 0 {
		return nil
	}

	close(p.done)

	p.mtx.Lock()
	defer p.mtx.Unlock()

	var issues []error
	for _, c := range p.conns {
		if err := c.Close(ctx); err != nil {
			issues = append(issues, err)
		}
	}

	if len(issues) > 0 {
		return errors.NewWithIssues("connection pool close failed", issues...)
	}

	return nil
}

func (p *pool) GetConn(e endpoint.Endpoint) Conn {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	if cc, ok := p.conns[e.Address()]; ok {
		return cc
	}
	cc := New(
		e,
		p.config,
		withOnClose(func(c Conn) {
			// conn.Conn.Close() must called on under locked p.mtx
			delete(p.conns, c.Endpoint().Address())
		}),
	)
	p.conns[e.Address()] = cc
	return cc
}

func (p *pool) connParker(ctx context.Context, interval time.Duration) {
	for {
		select {
		case <-p.done:
			return
		case <-time.After(interval):
			p.mtx.RLock()
			for _, c := range p.conns {
				select {
				case <-c.TTL():
					_ = c.Park(ctx)
				default:
					// pass
				}
			}
			p.mtx.RUnlock()
		}
	}
}

func NewPool(
	ctx context.Context,
	config Config,
) Pool {
	p := &pool{
		usages: 1,
		config: config,
		opts:   config.GrpcDialOptions(),
		conns:  make(map[string]Conn),
		done:   make(chan struct{}),
	}
	if ttl := config.ConnectionTTL(); ttl > 0 {
		go p.connParker(ctx, ttl/10)
	}
	return p
}

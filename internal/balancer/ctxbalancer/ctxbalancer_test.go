package ctxbalancer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/balancer/mock"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/conn"
)

func TestCtxBalancer_Create(t *testing.T) {
	b := Balancer([]conn.Conn{&mock.ConnMock{Address: "a1", NodeID: 1}, &mock.ConnMock{Address: "a2", NodeID: 2}}).(*ctxBalancer)
	checkOriginalBalancer := func() {
		require.Len(t, b.connMap, 2)
		require.Equal(t, b.connMap[1].(*mock.ConnMock).Address, "a1")
		require.Equal(t, b.connMap[2].(*mock.ConnMock).Address, "a2")
	}
	checkOriginalBalancer()

	b1 := b.Create([]conn.Conn{&mock.ConnMock{Address: "a3", NodeID: 3}, &mock.ConnMock{Address: "a4", NodeID: 4}, &mock.ConnMock{Address: "a5", NodeID: 5}}).(*ctxBalancer)
	checkOriginalBalancer()

	require.Len(t, b1.connMap, 3)
	require.Equal(t, b1.connMap[3].(*mock.ConnMock).Address, "a3")
	require.Equal(t, b1.connMap[4].(*mock.ConnMock).Address, "a4")
	require.Equal(t, b1.connMap[5].(*mock.ConnMock).Address, "a5")
}

func TestCtxBalancer_Next(t *testing.T) {
	b := Balancer([]conn.Conn{&mock.ConnMock{NodeID: 1, State: conn.Online}, &mock.ConnMock{NodeID: 2, State: conn.Banned}}).(*ctxBalancer)

	t.Run("EmptyContext", func(t *testing.T) {
		res := b.Next(context.Background(), true)
		require.Nil(t, res)
	})

	t.Run("WithPreferOnline", func(t *testing.T) {
		res := b.Next(WithEndpoint(context.Background(), &mock.EndpointMock{NodeIdField: 1}), true).(*mock.ConnMock)
		require.Equal(t, res.NodeID, uint32(1))
	})

	t.Run("WithPreferBanned", func(t *testing.T) {
		require.Equal(t, b.connMap[2].GetState(), conn.Banned)

		for _, allowBanned := range []bool{true, false} {
			res := b.Next(WithEndpoint(context.Background(), &mock.EndpointMock{NodeIdField: 2}), allowBanned).(*mock.ConnMock)
			require.Equal(t, res.NodeID, uint32(2))
		}
	})
}
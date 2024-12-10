package isolation

import (
	"database/sql"
	"database/sql/driver"
	"fmt"

	"github.com/ydb-platform/ydb-go-sdk/v3/internal/query/tx"
	"github.com/ydb-platform/ydb-go-sdk/v3/internal/xerrors"
	"github.com/ydb-platform/ydb-go-sdk/v3/query"
)

func ToYDB(opts driver.TxOptions) (txcControl tx.Option, err error) {
	level := sql.IsolationLevel(opts.Isolation)
	switch level {
	case sql.LevelDefault, sql.LevelSerializable:
		if !opts.ReadOnly {
			return query.WithSerializableReadWrite(), nil
		}
	case sql.LevelSnapshot:
		if opts.ReadOnly {
			return query.WithSnapshotReadOnly(), nil
		}
	}

	return nil, xerrors.WithStackTrace(fmt.Errorf(
		"unsupported transaction options: %+v", opts,
	))
}

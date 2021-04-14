package benchmark

import (
	"context"
	"database/sql"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
)

type benchDriver struct {
	db   *sql.DB
	mock sqlmock.Sqlmock
}

func NewBenchDriver() (*benchDriver, error) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	return &benchDriver{
		db:   db,
		mock: mock,
	}, err
}

type Tx struct {
	dialect.ExecQuerier
}

func (t Tx) Commit() error {
	return nil
}

func (t Tx) Rollback() error {
	return nil
}

func (b benchDriver) Exec(ctx context.Context, query string, args, v interface{}) error {
	argv, ok := args.([]interface{})
	if !ok {
		return fmt.Errorf("dialect/sql: invalid type %T. expect []interface{} for args", v)
	}
	switch v := v.(type) {
	case nil:
		if _, err := b.db.ExecContext(ctx, query, argv...); err != nil {
			return err
		}
	case *sql.Result:
		res, err := b.db.ExecContext(ctx, query, argv...)
		if err != nil {
			return err
		}
		*v = res
	default:
		return fmt.Errorf("dialect/sql: invalid type %T. expect *sql.Result", v)
	}
	return nil
}

func (b benchDriver) Query(ctx context.Context, query string, args, v interface{}) error {
	vr, ok := v.(*entsql.Rows)
	if !ok {
		return fmt.Errorf("dialect/sql: invalid type %T. expect *sql.Rows", v)
	}
	argv, ok := args.([]interface{})
	if !ok {
		return fmt.Errorf("dialect/sql: invalid type %T. expect []interface{} for args", args)
	}
	rows, err := b.db.QueryContext(ctx, query, argv...)
	if err != nil {
		return err
	}
	*vr = entsql.Rows{Rows: rows}
	return nil
}

func (b benchDriver) Tx(context.Context) (dialect.Tx, error) {
	return Tx{
		ExecQuerier: b,
	}, nil
}

func (b benchDriver) Close() error {
	return nil
}

func (b benchDriver) Dialect() string {
	return "mysql"
}

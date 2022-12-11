package adaptor

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// DBAdaptor コネクションを確立したデータストアを操作するためのメソッドをインターフェースとして提供
type DBAdaptor interface {
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	SelectForUpdate(dest interface{}, query string, tx *sqlx.Tx, args ...interface{}) error
	Queryx(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	NamedExec(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	Begin() *sqlx.Tx
}

type exdb struct {
	db *sqlx.DB
}

var _ sqlx.Ext = &exdb{}

func (s *exdb) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return s.db.Query(query, args...)
}

func (s *exdb) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	return s.db.Queryx(query, args...)
}

func (s *exdb) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	return s.db.QueryRowx(query, args...)
}

func (s *exdb) Exec(query string, args ...interface{}) (sql.Result, error) {
	return s.db.Exec(query, args...)
}

func (s *exdb) DriverName() string {
	return s.db.DriverName()
}

func (s *exdb) Rebind(query string) string {
	return s.db.Rebind(query)
}

func (s *exdb) BindNamed(query string, arg interface{}) (string, []interface{}, error) {
	return s.db.BindNamed(query, arg)
}

type basedb struct {
	sd *sqlx.DB
}

func (b *basedb) DB(ctx context.Context) *exdb {
	return &exdb{db: b.sd}
}

func (b *basedb) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	err := sqlx.Get(b.DB(ctx), dest, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (b *basedb) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	err := sqlx.Select(b.DB(ctx), dest, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (b *basedb) SelectForUpdate(dest interface{}, query string, tx *sqlx.Tx, args ...interface{}) error {
	err := tx.Select(dest, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (b *basedb) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	r, err := b.DB(ctx).Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (b *basedb) NamedExec(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	r, err := sqlx.NamedExec(b.DB(ctx), query, arg)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (b *basedb) Queryx(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	r, err := b.DB(ctx).Queryx(query, args...)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (b *basedb) Begin() *sqlx.Tx {
	tx := b.sd.MustBegin()
	return tx
}

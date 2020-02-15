package infra

import (
	"context"
	"database/sql"
)

// DB database manager
type DB interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	Open(driverName string, urlAddress string) error
	RawDB() (*sql.DB, error)
	Close() error
}

// Transactioner is the transaction interface for database
type Transactioner interface {
	Rollback() error
	Commit() error
	TxEnd(txFunc func() error) error
	TxBegin() (DB, error)
}

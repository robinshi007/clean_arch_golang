package database

import (

	// ignore package
	"context"
	"database/sql"
)

// DB database manager
type DB interface {
	ConnectDB(driverName string, urlAddress string) error
	CloseDB() error
	Begin(ctx context.Context) (context.Context, error)
	Rollback(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) (context.Context, error)
	Prepare(ctx context.Context, query string) (*sql.Stmt, error)
}

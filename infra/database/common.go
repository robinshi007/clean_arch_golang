package database

import (

	// ignore package
	"context"
	"database/sql"

	// pg and sqlite3
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

// DBM database manager
type DBM interface {
	ConnectDB(driverName string, urlAddress string) error
	CloseDB() error
	Begin(ctx context.Context) (context.Context, error)
	Rollback(ctx context.Context) (context.Context, error)
	Commit(ctx context.Context) (context.Context, error)
	Prepare(ctx context.Context, query string) (*sql.Stmt, error)
}

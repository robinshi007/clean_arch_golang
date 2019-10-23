package infra

import (
	"database/sql"
)

// DB database manager
type DB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	OpenDB(driverName string, urlAddress string) error
	Close() error
}

// Transactioner is the transaction interface for database
type Transactioner interface {
	Rollback() error
	Commit() error
	TxEnd(txFunc func() error) error
	TxBegin() (DB, error)
}

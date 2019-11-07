package database

import (
	"database/sql"

	// ignore package
	_ "github.com/lib/pq"

	"clean_arch/infra"
)

var dm infra.DB

// NewDB -
func NewDB(c *infra.Config) error {
	var dbma infra.DB
	dbma = &dbm{}
	err := dbma.Open(c.Database.DriverName, c.Database.URLAddress)
	if err != nil {
		return err
	}
	dm = dbma
	return nil
}

// GetDB get database manager
func GetDB() infra.DB {
	return dm
}

// dbm database manager
type dbm struct {
	DB *sql.DB
}

// ConnectDB database
func (m *dbm) Open(driverName, dataSourceName string) error {
	var err error
	m.DB, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		return err
	}
	return nil
}

// closeDB close database
func (m *dbm) Close() error {
	return m.DB.Close()
}

// Exec -
func (m *dbm) Exec(query string, args ...interface{}) (sql.Result, error) {
	return m.DB.Exec(query, args...)
}

// Prepare statement
func (m *dbm) Prepare(query string) (*sql.Stmt, error) {
	return m.DB.Prepare(query)
}

// Query -
func (m *dbm) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return m.DB.Query(query, args...)
}

// QueryRow
func (m *dbm) QueryRow(query string, args ...interface{}) *sql.Row {
	return m.DB.QueryRow(query, args...)
}

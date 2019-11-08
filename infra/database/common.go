package database

import (
	"database/sql"

	"clean_arch/infra"
)

// NewDB -
func NewDB(c *infra.Config) (infra.DB, error) {
	var dbma infra.DB
	dbma = &dbm{}
	err := dbma.Open(c.Database.DriverName, c.Database.URLAddress)
	if err != nil {
		return nil, err
	}
	return dbma, nil
}

// database manager
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

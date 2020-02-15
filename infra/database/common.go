package database

import (
	"github.com/jmoiron/sqlx"

	"clean_arch/infra"
)

const (
	// BeginTimeSQL - for log in sqlhooks
	BeginTimeSQL int = iota
)

// NewDB -
func NewDB(c *infra.Config) (infra.DB, error) {
	var dbma infra.DB

	if c.Database.DriverName == "postgres" {
		dbma = &pqsql{Mode: c.Mode}
	}
	err := dbma.Open(c.Database.DriverName, c.Database.URLAddress)
	if err != nil {
		return nil, err
	}
	return dbma, nil
}

// NewDBx -
func NewDBx(c *infra.Config) (*sqlx.DB, error) {
	dbma, err := NewDB(c)
	if err != nil {
		return nil, err
	}
	rawDb, err := dbma.RawDB()
	if err != nil {
		return nil, err
	}
	return sqlx.NewDb(rawDb, "postgres"), nil
}

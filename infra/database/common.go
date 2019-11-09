package database

import (
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
		dbma = &pqsql{}
	}
	err := dbma.Open(c.Database.DriverName, c.Database.URLAddress)
	if err != nil {
		return nil, err
	}
	return dbma, nil
}

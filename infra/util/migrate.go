package util

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	// drivers
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"clean_arch/infra"
)

// MigrationUp -
func MigrationUp(c *infra.Config, wd string) {
	db, err := sql.Open(c.Database.DriverName, c.Database.URLAddress)
	defer db.Close()
	FailedIf(err)

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	FailedIf(err)

	m, _ := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file:///%s/db/migrations", wd),
		c.Database.DriverName, driver)
	FailedIf(err)

	err = m.Up()
	if err != nil {
		if err.Error() == "no change" {
			fmt.Printf("[migration up] no change...")
		} else {
			FailedIf(err)
		}
	}
}

// MigrationDown -
func MigrationDown(c *infra.Config, wd string) {
	db, err := sql.Open(c.Database.DriverName, c.Database.URLAddress)
	defer db.Close()
	FailedIf(err)

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	FailedIf(err)

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file:///%s/db/migrations", wd),
		c.Database.DriverName, driver)
	FailedIf(err)

	err = m.Down()
	if err != nil {
		if err.Error() == "no change" {
			fmt.Printf("[migration down] no change...")
		} else {
			FailedIf(err)
		}
	}
}

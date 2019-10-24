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
	db, _ := sql.Open(c.Database.DriverName, c.Database.URLAddress)
	defer db.Close()
	driver, _ := postgres.WithInstance(db, &postgres.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file:///%s/db/migrations", wd),
		c.Database.DriverName, driver)
	m.Up()

}

// MigrationDown -
func MigrationDown(c *infra.Config, wd string) {
	db, _ := sql.Open(c.Database.DriverName, c.Database.URLAddress)
	defer db.Close()
	driver, _ := postgres.WithInstance(db, &postgres.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file:///%s/db/migrations", wd),
		c.Database.DriverName, driver)
	m.Down()

}

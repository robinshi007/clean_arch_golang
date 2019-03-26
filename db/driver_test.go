package db_test

import (
	"testing"

	"github.com/robinshi007/goweb/db"
)

func TestConnectDb(t *testing.T) {
	db, err := db.NewDb("localhost", "5432", "postgres", "postgres", "test")
	if err != nil {
		panic(err)
	}
	err = db.SQL.Ping()
	if err != nil {
		panic(err)
	}
}

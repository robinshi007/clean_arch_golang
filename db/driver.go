package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Db struct {
	SQL *sql.DB
	// Mgo *mgo.database
}

func NewDb(host, port, user, pass, dbname string) (*Db, error) {
	dbSource := ConnString(host, port, user, pass, dbname)
	db, err := sql.Open("postgres", dbSource)
	if err != nil {
		return nil, err
	}
	// check connection is good
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Db{SQL: db}, nil
}
func ConnString(host, port, user, pass, dbname string) string {
	return fmt.Sprintf(
		// "postgres://%s:%s@tcp(%s:%s)/%s?sslmode=disable",
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname,
	)
}

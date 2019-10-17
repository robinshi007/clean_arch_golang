package postgres_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	dba "clean_arch/adapter/database"
	"clean_arch/infra/database"
	"clean_arch/interface/postgres"
)

func getDBMock() (database.DBM, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}
	dbm := dba.NewDBMFromDB(db)
	return dbm, mock, err
}

func TestGetByID(t *testing.T) {
	db, mock, err := getDBMock()
	fmt.Println("dbm", db)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a sub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name", "description", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, "Hello", "Hello Desc", time.Now(), time.Now(), time.Now())

	query := "SELECT id, name, description, created_at, updated_at, deleted_at FROM users where id = $1"

	mock.ExpectQuery(query).WillReturnRows(rows)

	u := postgres.NewUserRepo(db)

	userID := int64(1)
	aUser, err := u.GetByID(context.TODO(), userID)
	assert.NoError(t, err)
	assert.NotNil(t, aUser)
}

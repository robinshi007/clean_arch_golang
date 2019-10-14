package postgres

import (
	"context"
	"database/sql"
	"time"

	dba "clean_arch/adapter/database"
	"clean_arch/infra/database"
)

var dbm *database.DBM

func init() {
	dbm = dba.GetDBM()
}

// begin begin transaction
func begin(ctx context.Context) (context.Context, error) {
	return (*dbm).Begin(ctx)
}

// rollback rollback transaction
func rollback(ctx context.Context) (context.Context, error) {
	return (*dbm).Rollback(ctx)
}

// commit commit transaction
func commit(ctx context.Context) (context.Context, error) {
	return (*dbm).Commit(ctx)
}

// prepare prepare statement
func prepare(ctx context.Context, query string) (*sql.Stmt, error) {
	return (*dbm).Prepare(ctx, query)
}

// TimeNow -
func TimeNow() time.Time {
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	return time.Now().In(location)
}

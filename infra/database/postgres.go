package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/gchaincl/sqlhooks"
	pq "github.com/lib/pq"

	"clean_arch/infra"
	"clean_arch/infra/util"
)

// database manager
type pqsql struct {
	DB   *sql.DB
	Mode string
}

var registerOnce sync.Once

// Open - ConnectDB database
func (p *pqsql) Open(driverName, dataSourceName string) error {
	var err error
	startupTimeout := 10 * time.Second
	startupDeadline := time.Now().Add(startupTimeout)

	// DB specific before check, e.g. force PostgreSQL session timezone to UTC.

	for {
		if time.Now().After(startupDeadline) {
			return fmt.Errorf("database did not start up in %s (%v)", startupTimeout, err)
		}
		err = p.openDBWithHooks(dataSourceName)
		if err == nil {
			err = p.DB.Ping()

		}
		if err != nil {
			time.Sleep(startupTimeout / 10)
			continue
		}

		// Database specific before check

		return nil
	}
}

func (p *pqsql) openDBWithHooks(dataSourceName string) error {
	registerOnce.Do(func() {
		sql.Register("postgres-proxy", sqlhooks.Wrap(&pq.Driver{}, &hook{
			Mode: p.Mode,
		}))
		fmt.Printf("PQSQL MODE: %s\n", p.Mode)
	})
	db, err := sql.Open("postgres-proxy", dataSourceName)
	p.DB = db
	return err
}

// closeDB close database
func (p *pqsql) Close() error {
	return p.DB.Close()
}

// Exec -
func (p *pqsql) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return p.DB.ExecContext(ctx, query, args...)
}

// Prepare statement
func (p *pqsql) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return p.DB.PrepareContext(ctx, query)
}

// QueryContext -
func (p *pqsql) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return p.DB.QueryContext(ctx, query, args...)
}

// QueryRow
func (p *pqsql) QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row {
	return p.DB.QueryRowContext(ctx, query, args...)
}

func (p *pqsql) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return p.DB.BeginTx(ctx, opts)
}

// Transaction

// DB doesn't rollback, do nothing here
func (p *pqsql) Rollback() error {
	return nil
}

//DB doesnt commit, do nothing here
func (p *pqsql) Commit() error {
	return nil
}

// TransactionBegin starts a transaction
func (p *pqsql) TxBegin() (infra.DB, error) {
	return nil, nil
}

// DB doesnt rollback, do nothing here
func (p *pqsql) TxEnd(txFunc func() error) error {
	return nil
}

type hook struct {
	Mode string
}

// Before implements sqlhooks.Hooks
func (h *hook) Before(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	// Print sql logs only in dev mode
	if h.Mode == "dev" {
		beginTime := time.Now()
		//fmt.Printf("> %s %q", query, args)
		util.CW(os.Stdout, util.Reset, "%s ", beginTime.Format(util.TimeFormatStr))
		util.CW(os.Stdout, util.NYellow, "\"%s %q\"", query, args)
		return context.WithValue(ctx, BeginTimeSQL, beginTime), nil
	}
	return ctx, nil
}

// After implements sqlhooks.Hooks
func (h *hook) After(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	// Print sql logs only in dev mode
	if h.Mode == "dev" {
		begin := ctx.Value(BeginTimeSQL).(time.Time)
		//fmt.Printf(". took: %s\n", time.Since(begin))
		util.CW(os.Stdout, util.Reset, " in ")
		timeDuration := time.Since(begin)
		if timeDuration > time.Millisecond*8 {
			util.CW(os.Stdout, util.NRed, "%s\n", timeDuration)
		} else {
			util.CW(os.Stdout, util.NGreen, "%s\n", timeDuration)
		}
	}
	return ctx, nil
}

// After implements sqlhooks.OnError
func (h *hook) OnError(ctx context.Context, err error, query string, args ...interface{}) error {
	return nil
}

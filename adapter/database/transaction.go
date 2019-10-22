package database

import (
	"clean_arch/infra/database"
)

// DB doesn't rollback, do nothing here
func (m *dbm) Rollback() error {
	return nil
}

//DB doesnt commit, do nothing here
func (m *dbm) Commit() error {
	return nil
}

// TransactionBegin starts a transaction
func (m *dbm) TxBegin() (database.DB, error) {
	return nil, nil
}

// DB doesnt rollback, do nothing here
func (m *dbm) TxEnd(txFunc func() error) error {
	return nil
}

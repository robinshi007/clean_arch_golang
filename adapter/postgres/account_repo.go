package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/keegancsmith/sqlf"

	"clean_arch/domain/model"
	"clean_arch/domain/repository"
	"clean_arch/infra/util"
	"clean_arch/registry"
)

// NewAccountRepo -
func NewAccountRepo() repository.AccountRepository {
	return &accountRepo{}
}

type accountRepo struct {
}

func (a *accountRepo) getBySQL(ctx context.Context, query string, args ...interface{}) ([]*model.UserAccount, error) {
	rows, err := registry.Db.QueryContext(ctx, "SELECT uid, name, email, password, created_at, updated_at FROM user_accounts "+query, args...)
	if err != nil {
		return nil, err
	}

	accounts := []*model.UserAccount{}
	defer rows.Close()
	for rows.Next() {
		//var deletedAt pq.NullTime
		var name sql.NullString
		account := model.UserAccount{}
		err := rows.Scan(
			&account.UID,
			&name,
			&account.Email,
			&account.Password,
			&account.CreatedAt,
			&account.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		if name.Valid {
			account.Name = name.String
		}
		//		if deletedAt.Valid {
		//			account.DeletedAt = deletedAt.Time
		//		}

		accounts = append(accounts, &account)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return accounts, nil
}

func (a *accountRepo) listSQL(opt repository.ListOptions) (conds []*sqlf.Query) {
	conds = []*sqlf.Query{}
	conds = append(conds, sqlf.Sprintf("deleted_at IS NULL"))
	if opt.Query != "" {
		query := "%" + opt.Query + "%"
		conds = append(conds, sqlf.Sprintf("name ILIKE %s", query))
	}
	return conds
}

func (a *accountRepo) FindAll(ctx context.Context, opt *repository.ListOptions) ([]*model.UserAccount, error) {
	if opt == nil {
		opt = &repository.ListOptions{}
	}
	conds := a.listSQL(*opt)
	q := sqlf.Sprintf("WHERE %s ORDER BY uid ASC %s", sqlf.Join(conds, "AND"), opt.LimitOffset.SQL())
	return a.getBySQL(ctx, q.Query(sqlf.PostgresBindVar), q.Args()...)
}

func (a *accountRepo) FindByID(ctx context.Context, id int64) (*model.UserAccount, error) {
	rows, err := a.getBySQL(ctx, "WHERE deleted_at IS NULL AND uid=$1 LIMIT 1", strconv.FormatInt(id, 10))
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, model.ErrEntityNotFound
	}
	return rows[0], nil
}
func (a *accountRepo) FindByEmail(ctx context.Context, email string) (*model.UserAccount, error) {
	rows, err := a.getBySQL(ctx, "WHERE deleted_at IS NULL AND email=$1 LIMIT 1", email)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, model.ErrEntityNotFound
	}
	return rows[0], nil
}
func (a *accountRepo) FindByName(ctx context.Context, name string) (*model.UserAccount, error) {
	rows, err := a.getBySQL(ctx, "WHERE deleted_at IS NULL AND name=$1 LIMIT 1", name)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, model.ErrEntityNotFound
	}
	return rows[0], nil
}
func (a *accountRepo) Create(ctx context.Context, account *model.UserAccount) (int64, error) {
	query := "INSERT INTO user_accounts (email, name, password, password_hash_argorithm, created_at,updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING uid"
	tx, err := registry.Db.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}
	defer func() {
		if err != nil {
			util.CW(os.Stdout, util.NRed, "\"%s\"\n", err.Error())
			rollErr := tx.Rollback()
			if rollErr != nil {
				fmt.Println("rollback error:", rollErr.Error())
			}
		}
	}()
	account.CreatedAt = time.Now()
	account.UpdatedAt = account.CreatedAt

	err = tx.QueryRowContext(
		ctx,
		query,
		account.Email,
		account.Name,
		account.Password,
		account.PasswordHashArgorithm,
		account.CreatedAt,
		account.UpdatedAt,
	).Scan(&account.UID)
	if err != nil {
		return -1, err
	}
	// create user profile
	queryProfile := "INSERT INTO user_profiles (uid, email, created_at, updated_at) VALUES ($1, $2, $3, $4)"
	profile := &model.UserProfile{
		UID:   account.UID,
		Email: account.Email,
	}
	profile.CreatedAt = time.Now()
	profile.UpdatedAt = profile.CreatedAt

	_, err = tx.ExecContext(
		ctx,
		queryProfile,
		profile.UID,
		profile.Email,
		profile.CreatedAt,
		profile.UpdatedAt,
	)
	if err != nil {
		return -1, err
	}
	err = tx.Commit()
	if err != nil {
		return -1, err
	}
	//fmt.Println("uid", account.UID)
	return account.UID, nil
}
func (a *accountRepo) Update(ctx context.Context, account *model.UserAccount) (*model.UserAccount, error) {
	tx, err := registry.Db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			util.CW(os.Stdout, util.NRed, "\"%s\"\n", err.Error())
			rollErr := tx.Rollback()
			if rollErr != nil {
				fmt.Println("rollback error:", rollErr.Error())
			}
			return
		}
		err = tx.Commit()
	}()

	query := "UPDATE user_accounts SET name=$1, updated_at=$2 WHERE uid=$3 AND deleted_at IS NULL"
	account.UpdatedAt = time.Now()
	_, err = tx.ExecContext(ctx, query, account.Name, account.UpdatedAt, account.UID)
	if err != nil {
		return HandleAccountPqErr(err)
	}
	return account, nil
}
func (a *accountRepo) UpdatePassword(ctx context.Context, account *model.UserAccount) (*model.UserAccount, error) {
	tx, err := registry.Db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			util.CW(os.Stdout, util.NRed, "\"%s\"\n", err.Error())
			rollErr := tx.Rollback()
			if rollErr != nil {
				fmt.Println("rollback error:", rollErr.Error())
			}
			return
		}
		err = tx.Commit()
	}()

	query := "UPDATE user_accounts SET password=$1, updated_at=$2 WHERE uid=$3 AND deleted_at IS NULL"
	account.UpdatedAt = time.Now()
	_, err = tx.ExecContext(ctx, query, account.Password, account.UpdatedAt, account.UID)
	if err != nil {
		return HandleAccountPqErr(err)
	}
	return account, nil
}
func (a *accountRepo) Delete(ctx context.Context, id int64) error {
	timeNow := time.Now()
	query := "UPDATE user_accounts SET updated_at=$1, deleted_at=$2 WHERE uid=$3"
	res, err := registry.Db.ExecContext(ctx, query, timeNow, timeNow, strconv.FormatInt(id, 10))
	if err != nil {
		util.CW(os.Stdout, util.NRed, "\"%s\"\n", err.Error())
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return model.ErrEntityNotFound
	}
	queryProfile := "UPDATE user_profiles SET updated_at=$1, deleted_at=$2 WHERE uid=$3"
	res, err = registry.Db.ExecContext(ctx, queryProfile, timeNow, timeNow, strconv.FormatInt(id, 10))
	if err != nil {
		util.CW(os.Stdout, util.NRed, "\"%s\"\n", err.Error())
		return err
	}
	count, err = res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return model.ErrEntityNotFound
	}
	return nil
}
func (a *accountRepo) ExistsByEmail(ctx context.Context, email string) error {
	account, err := a.FindByEmail(ctx, email)
	if account != nil {
		return model.ErrEntityUniqueConflict
	}
	if err != nil {
		return err
	}
	return nil
}
func (a *accountRepo) ExistsByName(ctx context.Context, name string) error {
	account, err := a.FindByName(ctx, name)
	if account != nil {
		return model.ErrEntityUniqueConflict
	}
	if err != nil {
		return err
	}
	return nil
}

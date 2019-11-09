package postgres

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/keegancsmith/sqlf"

	"clean_arch/domain/model"
	"clean_arch/domain/repository"
	"clean_arch/infra"
	"clean_arch/infra/util"
	"clean_arch/registry"
)

// NewUserRepo -
func NewUserRepo(conn infra.DB) repository.UserRepository {
	return &userRepo{
		DB: conn,
	}
}

type userRepo struct {
	DB infra.DB
}

func (u *userRepo) getBySQL(ctx context.Context, query string, args ...interface{}) ([]*model.User, error) {
	rows, err := registry.Db.QueryContext(ctx, "SELECT id, name, description, created_at, updated_at FROM users "+query, args...)
	if err != nil {
		return nil, err
	}

	users := []*model.User{}
	defer rows.Close()
	for rows.Next() {
		//var deletedAt pq.NullTime
		user := model.User{}
		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Description,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		//		if deletedAt.Valid {
		//			user.DeletedAt = deletedAt.Time
		//		}

		users = append(users, &user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
func (u *userRepo) listSQL(opt repository.UserListOptions) (conds []*sqlf.Query) {
	conds = []*sqlf.Query{}
	conds = append(conds, sqlf.Sprintf("deleted_at IS NULL"))
	if opt.Query != "" {
		query := "%" + opt.Query + "%"
		conds = append(conds, sqlf.Sprintf("name ILIKE %s OR discription ILIKE %s", query, query))
	}
	return conds
}

// GetAll -
func (u *userRepo) GetAll(ctx context.Context, opt *repository.UserListOptions) ([]*model.User, error) {
	if opt == nil {
		opt = &repository.UserListOptions{}
	}
	conds := u.listSQL(*opt)
	q := sqlf.Sprintf("WHERE %s ORDER BY id ASC %s", sqlf.Join(conds, "AND"), opt.LimitOffset.SQL())
	return u.getBySQL(ctx, q.Query(sqlf.PostgresBindVar), q.Args()...)
}

// GetByID -
func (u *userRepo) GetByID(ctx context.Context, id int64) (*model.User, error) {
	rows, err := u.getBySQL(ctx, "WHERE deleted_at IS NULL AND id=$1 LIMIT 1", strconv.FormatInt(id, 10))
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, model.ErrEntityNotFound
	}
	return rows[0], nil
}

// GetByName -
func (u *userRepo) GetByName(ctx context.Context, name string) (*model.User, error) {
	rows, err := u.getBySQL(ctx, "WHERE deleted_at IS NULL AND name=$1 LIMIT 1", name)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, model.ErrEntityNotFound
	}
	return rows[0], nil
}

// Create -
func (u *userRepo) Create(ctx context.Context, user *model.User) (int64, error) {
	query := "INSERT INTO users (name,description,created_at,updated_at) VALUES ($1, $2, $3, $4) RETURNING id"
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
			return
		}
		err = tx.Commit()
	}()
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt

	err = tx.QueryRowContext(
		ctx,
		query,
		user.Name,
		user.Description,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID)
	if err != nil {
		return -1, err
	}
	return user.ID, nil
}

// Update -
func (u *userRepo) Update(ctx context.Context, user *model.User) (*model.User, error) {
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

	query := "UPDATE users SET name=$1, description=$2, updated_at=$3 WHERE id=$4 AND deleted_at IS NULL RETURNING id"
	user.UpdatedAt = time.Now()
	_, err = tx.ExecContext(ctx, query, user.Name, user.Description, user.UpdatedAt, user.ID)
	if err != nil {
		return HandleUserPqErr(err)
	}
	return user, nil
}

// Delete -
func (u *userRepo) Delete(ctx context.Context, id int64) error {
	query := "UPDATE users SET updated_at=$1, deleted_at=$2 WHERE id=$3"
	timeNow := time.Now()
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
	return nil
}

// DuplicatedByName -
func (u *userRepo) DuplicatedByName(ctx context.Context, name string) error {
	user, err := u.GetByName(ctx, name)
	if user != nil {
		return model.ErrEntityUniqueConflict
	}
	if err != nil {
		return err
	}
	return nil
}

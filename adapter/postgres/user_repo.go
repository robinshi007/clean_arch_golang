package postgres

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/keegancsmith/sqlf"

	"clean_arch/domain/model"
	"clean_arch/domain/repository"
	"clean_arch/infra/util"
	"clean_arch/registry"
)

// NewUserRepo -
func NewUserRepo() repository.UserRepository {
	return &userRepo{}
}

type userRepo struct {
}

func (u *userRepo) getBySQL(ctx context.Context, query string, args ...interface{}) ([]*model.User, error) {
	users := []*model.User{}
	err := registry.Db.SelectContext(ctx, &users, "SELECT id, name, description, created_at, updated_at FROM users "+query, args...)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (u *userRepo) listSQL(opt repository.ListOptions) (conds []*sqlf.Query) {
	conds = []*sqlf.Query{}
	conds = append(conds, sqlf.Sprintf("deleted_at IS NULL"))
	if opt.Query != "" {
		query := "%" + opt.Query + "%"
		conds = append(conds, sqlf.Sprintf("name ILIKE %s OR discription ILIKE %s", query, query))
	}
	return conds
}

// Count -
func (u *userRepo) Count(ctx context.Context) (int64, error) {
	var count int64
	row := registry.Db.QueryRowContext(ctx, "SELECT Count(*) from users WHERE deleted_at IS NULL")
	err := row.Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

// FindAll -
func (u *userRepo) FindAll(ctx context.Context, opt *repository.ListOptions) ([]*model.User, error) {
	if opt == nil {
		opt = &repository.ListOptions{}
	}
	conds := u.listSQL(*opt)
	q := sqlf.Sprintf("WHERE %s ORDER BY id ASC %s", sqlf.Join(conds, "AND"), opt.LimitOffset.SQL())
	return u.getBySQL(ctx, q.Query(sqlf.PostgresBindVar), q.Args()...)
}

// FindByID -
func (u *userRepo) FindByID(ctx context.Context, id int64) (*model.User, error) {
	rows, err := u.getBySQL(ctx, "WHERE deleted_at IS NULL AND id=$1 LIMIT 1", strconv.FormatInt(id, 10))
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, model.ErrEntityNotFound
	}
	return rows[0], nil
}

// FindByName -
func (u *userRepo) FindByName(ctx context.Context, name string) (*model.User, error) {
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
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt

	err := registry.Db.QueryRowContext(
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
	query := "UPDATE users SET name=$1, description=$2, updated_at=$3 WHERE id=$4 AND deleted_at IS NULL RETURNING id"
	user.UpdatedAt = time.Now()
	res, err := registry.Db.ExecContext(ctx, query, user.Name, user.Description, user.UpdatedAt, user.ID)
	if err != nil {
		return HandleUserPqErr(err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, model.ErrEntityNotFound
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

// ExistsByName -
func (u *userRepo) ExistsByName(ctx context.Context, name string) error {
	user, err := u.FindByName(ctx, name)
	if user != nil {
		return model.ErrEntityUniqueConflict
	}
	if err != nil {
		return err
	}
	return nil
}

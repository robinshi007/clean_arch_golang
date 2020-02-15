package postgres

import (
	"context"
	"strconv"

	"github.com/keegancsmith/sqlf"

	"clean_arch/domain/model"
	"clean_arch/domain/repository"
	"clean_arch/registry"
)

// NewProfileRepo -
func NewProfileRepo() repository.ProfileRepository {
	return &profileRepo{}
}

type profileRepo struct {
}

func (pr *profileRepo) getBySQL(ctx context.Context, query string, args ...interface{}) ([]*model.UserProfile, error) {
	profiles := []*model.UserProfile{}
	err := registry.Db.SelectContext(ctx, &profiles, "SELECT uid, email, created_at, updated_at FROM user_profiles "+query, args...)
	if err != nil {
		return nil, err
	}
	return profiles, nil
}

func (pr *profileRepo) listSQL(opt repository.ListOptions) (conds []*sqlf.Query) {
	conds = []*sqlf.Query{}
	conds = append(conds, sqlf.Sprintf("deleted_at IS NULL"))
	if opt.Query != "" {
		query := "%" + opt.Query + "%"
		conds = append(conds, sqlf.Sprintf("name ILIKE %s", query))
	}
	return conds
}

func (pr *profileRepo) FindAll(ctx context.Context, opt *repository.ListOptions) ([]*model.UserProfile, error) {
	if opt == nil {
		opt = &repository.ListOptions{}
	}
	conds := pr.listSQL(*opt)
	q := sqlf.Sprintf("WHERE %s ORDER BY uid ASC %s", sqlf.Join(conds, "AND"), opt.LimitOffset.SQL())
	return pr.getBySQL(ctx, q.Query(sqlf.PostgresBindVar), q.Args()...)
}

func (pr *profileRepo) FindByID(ctx context.Context, id int64) (*model.UserProfile, error) {
	rows, err := pr.getBySQL(ctx, "WHERE deleted_at IS NULL AND uid=$1 LIMIT 1", strconv.FormatInt(id, 10))
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, model.ErrEntityNotFound
	}
	return rows[0], nil
}
func (pr *profileRepo) FindByEmail(ctx context.Context, email string) (*model.UserProfile, error) {
	rows, err := pr.getBySQL(ctx, "WHERE deleted_at IS NULL AND email=$1 LIMIT 1", email)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, model.ErrEntityNotFound
	}
	return rows[0], nil
}

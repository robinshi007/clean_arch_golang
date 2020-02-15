package postgres

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"clean_arch/domain/model"
	"clean_arch/domain/repository"
	"clean_arch/infra/util"
	"clean_arch/registry"

	"github.com/keegancsmith/sqlf"
)

// NewRedirectRepo -
func NewRedirectRepo() repository.RedirectRepository {
	return &redirectRepo{}
}

type redirectRepo struct {
}

func (r *redirectRepo) getBySQL(ctx context.Context, query string, args ...interface{}) ([]*model.Redirect, error) {
	redirects := []*model.Redirect{}
	rows, err := registry.Db.QueryxContext(ctx, `SELECT r.id, r.code, r.url, r.created_at, 
up.uid "created_by.uid", 
up.email "created_by.email", 
up.created_at "created_by.created_at", 
up.updated_at "created_by.updated_at"
FROM redirects r `+query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var r model.Redirect
		err = rows.StructScan(&r)
		if err != nil {
			return nil, err
		}
		redirects = append(redirects, &r)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return redirects, nil
}

func (r *redirectRepo) listSQL(opt repository.ListOptions) (conds []*sqlf.Query) {
	conds = []*sqlf.Query{}
	conds = append(conds, sqlf.Sprintf("r.deleted_at IS NULL"))
	return conds
}

// Count -
func (r *redirectRepo) Count(ctx context.Context) (int64, error) {
	var count int64
	row := registry.Db.QueryRowContext(ctx, "SELECT Count(*) from redirects WHERE deleted_at IS NULL")
	err := row.Scan(&count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (r *redirectRepo) FindAll(ctx context.Context, opt *repository.ListOptions) ([]*model.Redirect, error) {
	if opt == nil {
		opt = &repository.ListOptions{}
	}
	conds := r.listSQL(*opt)
	q := sqlf.Sprintf("JOIN user_profiles up ON r.created_by_id = up.uid WHERE %s ORDER BY id ASC %s", sqlf.Join(conds, "AND"), opt.LimitOffset.SQL())
	return r.getBySQL(ctx, q.Query(sqlf.PostgresBindVar), q.Args()...)
}

// FindByID -
func (r *redirectRepo) FindByID(ctx context.Context, id int64) (*model.Redirect, error) {
	rows, err := r.getBySQL(ctx, "JOIN user_profiles up ON r.created_by_id = up.uid WHERE r.deleted_at IS NULL AND r.id=$1 LIMIT 1", strconv.FormatInt(id, 10))
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, model.ErrEntityNotFound
	}
	return rows[0], nil
}

// FindByCode -
func (r *redirectRepo) FindByCode(ctx context.Context, code string) (*model.Redirect, error) {
	rows, err := r.getBySQL(ctx, "JOIN user_profiles up ON r.created_by_id = up.uid WHERE r.deleted_at IS NULL AND r.code=$1 LIMIT 1", code)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, model.ErrEntityNotFound
	}
	return rows[0], nil
}

// FindByURL -
func (r *redirectRepo) FindByURL(ctx context.Context, code string) (*model.Redirect, error) {
	rows, err := r.getBySQL(ctx, "JOIN user_profiles up ON r.created_by_id = up.uid WHERE r.deleted_at IS NULL AND r.url=$1 LIMIT 1", code)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, model.ErrEntityNotFound
	}
	return rows[0], nil
}

// Create -
func (r *redirectRepo) Create(ctx context.Context, redirect *model.Redirect) (int64, error) {
	query := "INSERT INTO redirects (code,url,created_by_id,created_at) VALUES ($1, $2, $3, $4) RETURNING id"
	redirect.CreatedAt = time.Now()
	err := registry.Db.QueryRowContext(
		ctx,
		query,
		redirect.Code,
		redirect.URL,
		redirect.CreatedBy.UID,
		redirect.CreatedAt,
	).Scan(&redirect.ID)
	if err != nil {
		fmt.Println("Redirect.Create:", err)
		return -1, err
	}
	return redirect.ID, nil
}

// Delete -
func (r *redirectRepo) Delete(ctx context.Context, id int64) error {
	query := "UPDATE redirects SET deleted_at=$1 WHERE id=$2"
	timeNow := time.Now()
	res, err := registry.Db.ExecContext(ctx, query, timeNow, strconv.FormatInt(id, 10))
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

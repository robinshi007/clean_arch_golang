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
	rows, err := registry.Db.QueryContext(ctx, "SELECT id, code, url, created_at FROM redirects "+query, args...)
	if err != nil {
		return nil, err
	}

	redirects := []*model.Redirect{}
	defer rows.Close()
	for rows.Next() {
		//var deletedAt pq.NullTime
		redirect := model.Redirect{}
		err := rows.Scan(
			&redirect.ID,
			&redirect.Code,
			&redirect.URL,
			&redirect.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		//		if deletedAt.Valid {
		//			redirect.DeletedAt = deletedAt.Time
		//		}

		redirects = append(redirects, &redirect)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return redirects, nil
}

func (r *redirectRepo) listSQL(opt repository.RedirectListOptions) (conds []*sqlf.Query) {
	conds = []*sqlf.Query{}
	conds = append(conds, sqlf.Sprintf("deleted_at IS NULL"))
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

func (r *redirectRepo) FindAll(ctx context.Context, opt *repository.RedirectListOptions) ([]*model.Redirect, error) {
	if opt == nil {
		opt = &repository.RedirectListOptions{}
	}
	conds := r.listSQL(*opt)
	q := sqlf.Sprintf("WHERE %s ORDER BY id ASC %s", sqlf.Join(conds, "AND"), opt.LimitOffset.SQL())
	return r.getBySQL(ctx, q.Query(sqlf.PostgresBindVar), q.Args()...)
}

// FindByID -
func (r *redirectRepo) FindByID(ctx context.Context, id int64) (*model.Redirect, error) {
	rows, err := r.getBySQL(ctx, "WHERE deleted_at IS NULL AND id=$1 LIMIT 1", strconv.FormatInt(id, 10))
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
	rows, err := r.getBySQL(ctx, "WHERE deleted_at IS NULL AND code=$1 LIMIT 1", code)
	if err != nil {
		return nil, err
	}
	if len(rows) == 0 {
		return nil, model.ErrEntityNotFound
	}
	return rows[0], nil
}

// Create -
func (r *redirectRepo) Save(ctx context.Context, redirect *model.Redirect) (int64, error) {
	query := "INSERT INTO redirects (code,url,created_at) VALUES ($1, $2, $3) RETURNING id"
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
	redirect.CreatedAt = time.Now()

	err = tx.QueryRowContext(
		ctx,
		query,
		redirect.Code,
		redirect.URL,
		redirect.CreatedAt,
	).Scan(&redirect.ID)
	if err != nil {
		return -1, err
	}
	return redirect.ID, nil
}

// Delete -
func (u *redirectRepo) Delete(ctx context.Context, id int64) error {
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

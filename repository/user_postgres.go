package repository

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	pq "github.com/lib/pq"

	"github.com/robinshi007/goweb/db"
	model "github.com/robinshi007/goweb/model"
)

func NewUserRepo(conn *db.Db) UserRepo {
	return &postgresUserRepo{
		Conn: conn.SQL,
	}
}

type postgresUserRepo struct {
	Conn *sql.DB
}

func (p *postgresUserRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*model.User, error) {
	result := make([]*model.User, 0)
	rows, err := p.Conn.Query(query)
	defer rows.Close()
	if err != nil {
		return result, err
	}
	for rows.Next() {
		var deletedAt pq.NullTime
		data := new(model.User)
		err := rows.Scan(
			&data.Id,
			&data.Name,
			&data.Desc,
			&data.CreatedAt,
			&data.UpdatedAt,
			&deletedAt,
		)
		if err != nil {
			return result, err
		}
		if deletedAt.Valid {
			data.DeletedAt = deletedAt.Time
		}
		result = append(result, data)
	}
	return result, nil
}
func (p *postgresUserRepo) Fetch(ctx context.Context, num int64) ([]*model.User, error) {
	query := "SELECT id, name, description, created_at, updated_at, deleted_at FROM users LIMIT " + strconv.FormatInt(num, 10)
	return p.fetch(ctx, query, num)
}

func (p *postgresUserRepo) GetById(ctx context.Context, id int64) (*model.User, error) {
	query := "SELECT id, name, description, created_at, updated_at, deleted_at FROM users where id = " + strconv.FormatInt(id, 10) + ";"
	rows, err := p.fetch(ctx, query)
	if err != nil {
		return nil, err
	}
	payload := &model.User{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, model.ErrNotFound
	}
	return payload, nil
}

func (p *postgresUserRepo) Create(ctx context.Context, u *model.User) (int64, error) {
	var userId int
	query := "INSERT INTO users (name,description,created_at,updated_at) VALUES ($1, $2, $3, $4) RETURNING id"
	err := p.Conn.QueryRow(query, u.Name, u.Desc, time.Now(), time.Now()).Scan(&userId)
	if err != nil {
		return -1, err
	}
	return int64(userId), err
}

func (p *postgresUserRepo) Update(ctx context.Context, u *model.User) (*model.User, error) {
	query := "UPDATE users SET name=$1, description=$2, updated_at=$3 where id=$4 RETURNING id"

	u.UpdatedAt = time.Now()
	_, err := p.Conn.Exec(query, u.Name, u.Desc, u.UpdatedAt, u.Id)
	if err != nil {
		return nil, err
	}
	return u, nil
}
func (p *postgresUserRepo) Delete(ctx context.Context, id int64) (bool, error) {
	query := "DELETE FROM users where id=$1"
	_, err := p.Conn.Exec(query, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

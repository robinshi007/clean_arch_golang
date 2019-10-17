package postgres

import (
	"context"
	"fmt"
	"strconv"
	"time"

	pq "github.com/lib/pq"

	"clean_arch/domain/model"
	"clean_arch/domain/repository"
	"clean_arch/infra/database"
)

// NewUserRepo -
func NewUserRepo(conn database.DBM) repository.UserRepository {
	return &postgresUserRepo{
		DBM: conn,
	}
}

type postgresUserRepo struct {
	DBM database.DBM
}

func (p *postgresUserRepo) fetch(c context.Context, query string, args ...interface{}) ([]*model.User, error) {
	stmt, err := p.DBM.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(c)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	result := make([]*model.User, 0)
	for rows.Next() {
		var deletedAt pq.NullTime
		data := new(model.User)
		err := rows.Scan(
			&data.ID,
			&data.Name,
			&data.Description,
			&data.CreatedAt,
			&data.UpdatedAt,
			&deletedAt,
		)
		fmt.Println("row", data)
		if err != nil {
			return result, err
		}
		if deletedAt.Valid {
			data.DeletedAt = deletedAt.Time
		}
		result = append(result, data)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
func (p *postgresUserRepo) Fetch(c context.Context, num int64) ([]*model.User, error) {
	query := "SELECT id, name, description, created_at, updated_at, deleted_at FROM users LIMIT " + strconv.FormatInt(num, 10)
	return p.fetch(c, query, num)
}

func (p *postgresUserRepo) GetByID(c context.Context, id int64) (*model.User, error) {
	query := "SELECT id, name, description, created_at, updated_at, deleted_at FROM users where id = " + strconv.FormatInt(id, 10) + ";"
	rows, err := p.fetch(c, query)
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
func (p *postgresUserRepo) GetByName(c context.Context, name string) (*model.User, error) {
	query := "SELECT id, name, description, created_at, updated_at, deleted_at FROM users where name = $1;"
	rows, err := p.fetch(c, query, name)
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

func (p *postgresUserRepo) Create(c context.Context, u *model.User) (int64, error) {
	var userID int
	now := time.Now()
	query := "INSERT INTO users (name,description,created_at,updated_at) VALUES ($1, $2, $3, $4) RETURNING id"

	stmt, err := p.DBM.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return -1, err
	}
	err = stmt.QueryRowContext(c, u.Name, u.Description, now, now).Scan(&userID)
	if err != nil {
		return -1, err
	}
	return int64(userID), err
}

func (p *postgresUserRepo) Update(c context.Context, u *model.User) (*model.User, error) {
	query := "UPDATE users SET name=$1, description=$2, updated_at=$3 where id=$4 RETURNING id"

	u.UpdatedAt = time.Now()
	stmt, err := p.DBM.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return HandleUserPqErr(err)
	}
	_, err = stmt.ExecContext(c, u.Name, u.Description, u.UpdatedAt, u.ID)
	if err != nil {
		return HandleUserPqErr(err)
	}
	return u, nil
}
func (p *postgresUserRepo) Delete(c context.Context, id int64) error {
	query := "DELETE FROM users where id=$1"

	stmt, err := p.DBM.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(c, id)
	if err != nil {
		return err
	}
	return nil
}

package postgres

import (
	"context"
	"strconv"
	"time"

	pq "github.com/lib/pq"

	"clean_arch/domain/model"
	"clean_arch/domain/repository"
	"clean_arch/infra/database"
)

// NewUserRepo -
func NewUserRepo(conn database.DB) repository.UserRepository {
	return &postgresUserRepo{
		DB: conn,
	}
}

type postgresUserRepo struct {
	DB database.DB
}

func (p *postgresUserRepo) fetch(query string, args ...interface{}) ([]*model.User, error) {
	result := make([]*model.User, 0)
	ctx := context.Background()
	stmt, err := p.DB.Prepare(ctx, query)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
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
func (p *postgresUserRepo) Fetch(num int64) ([]*model.User, error) {
	query := "SELECT id, name, description, created_at, updated_at, deleted_at FROM users LIMIT " + strconv.FormatInt(num, 10)
	return p.fetch(query, num)
}

func (p *postgresUserRepo) GetByID(id int64) (*model.User, error) {
	query := "SELECT id, name, description, created_at, updated_at, deleted_at FROM users where id = " + strconv.FormatInt(id, 10) + ";"
	rows, err := p.fetch(query)
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
func (p *postgresUserRepo) GetByName(name string) (*model.User, error) {
	return &model.User{Name: "haha"}, nil
}

func (p *postgresUserRepo) Create(u *model.User) (int64, error) {
	var userID int
	now := time.Now()
	query := "INSERT INTO users (name,description,created_at,updated_at) VALUES ($1, $2, $3, $4) RETURNING id"

	ctx := context.Background()
	stmt, err := p.DB.Prepare(ctx, query)
	defer stmt.Close()
	if err != nil {
		return -1, err
	}
	err = stmt.QueryRowContext(ctx, u.Name, u.Description, now, now).Scan(&userID)
	if err != nil {
		return -1, err
	}
	return int64(userID), err
}

func (p *postgresUserRepo) Update(u *model.User) (*model.User, error) {
	query := "UPDATE users SET name=$1, description=$2, updated_at=$3 where id=$4 RETURNING id"

	u.UpdatedAt = time.Now()

	ctx := context.Background()
	stmt, err := p.DB.Prepare(ctx, query)
	defer stmt.Close()
	if err != nil {
		return HandlePqErr(err)
	}
	_, err = stmt.ExecContext(ctx, u.Name, u.Description, u.UpdatedAt, u.ID)
	if err != nil {
		return HandlePqErr(err)
	}
	return u, nil
}
func (p *postgresUserRepo) Delete(id int64) (bool, error) {
	query := "DELETE FROM users where id=$1"

	ctx := context.Background()
	stmt, err := p.DB.Prepare(ctx, query)
	defer stmt.Close()
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}

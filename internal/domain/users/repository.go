package users

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lucianboboc/todo-api/internal/intrastructure/database"
)

type Repository interface {
	Create(ctx context.Context, user *User) error
	GetUsers(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id int) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}

func NewRepository(db database.Database) Repository {
	return repository{
		DB: db,
	}
}

type repository struct {
	DB database.Database
}

func (r repository) Create(ctx context.Context, user *User) error {
	query := `
	INSERT INTO users (first_name, last_name, email, password_hash)
	VALUES ($1, $2, $3, $4)
	RETURNING id, created_at
`

	args := []any{user.FirstName, user.LastName, user.Email, user.PasswordHash}
	err := r.DB.QueryRowContext(ctx, query, args...).Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		if database.IsUniqueConstraint(err) {
			return database.ErrRecordAlreadyExists
		}
		return err
	}
	return nil
}

func (r repository) GetUsers(ctx context.Context) ([]User, error) {
	query := `
	SELECT id, first_name, last_name, email, password_hash, created_at FROM users
`
	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := make([]User, 0)
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.PasswordHash,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return res, nil
}

func (r repository) GetByID(ctx context.Context, id int) (*User, error) {
	query := `
	SELECT id, first_name, last_name, email, password_hash, created_at FROM users WHERE id = $1
`
	var user User
	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, database.ErrNoRecordsFound
		}
		return nil, err
	}
	return &user, nil
}

func (r repository) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := `
	SELECT id, first_name, last_name, email, password_hash, created_at FROM users WHERE email = $1
`
	var user User
	err := r.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, database.ErrNoRecordsFound
		}
		return nil, err
	}
	return &user, nil
}

package users

import (
	"context"
	"database/sql"
)

type Repository interface {
	Create(ctx context.Context, user *User) error
	GetByID(ctx context.Context, id int) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}

func NewRepository(db *sql.DB) Repository {
	return repository{
		DB: db,
	}
}

type repository struct {
	DB *sql.DB
}

// TODO: Implement database operations
func (r repository) Create(ctx context.Context, user *User) error {
	return nil
}

func (r repository) GetByID(ctx context.Context, id int) (*User, error) {
	return nil, nil
}

func (r repository) GetByEmail(ctx context.Context, email string) (*User, error) {
	return nil, nil
}

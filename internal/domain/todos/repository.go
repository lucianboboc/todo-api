package todos

import (
	"context"
	"database/sql"
)

type Repository interface {
	Create(ctx context.Context, todo Todo) error
	GetById(ctx context.Context, id int) (*Todo, error)
	GetAll(ctx context.Context) ([]Todo, error)
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
func (r repository) Create(ctx context.Context, todo Todo) error {
	return nil
}

func (r repository) GetById(ctx context.Context, id int) (*Todo, error) {
	return nil, nil
}

func (r repository) GetAll(ctx context.Context) ([]Todo, error) {
	return nil, nil
}

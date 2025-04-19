package todos

import (
	"context"
	"github.com/lucianboboc/todo-api/internal/intrastructure/database"
)

type Repository interface {
	Create(ctx context.Context, todo *Todo) error
	GetById(ctx context.Context, id int) (*Todo, error)
	GetAll(ctx context.Context) ([]Todo, error)
}

func NewRepository(db database.Database) Repository {
	return repository{
		DB: db,
	}
}

type repository struct {
	DB database.Database
}

func (r repository) Create(ctx context.Context, todo *Todo) error {
	query := `
	INSERT INTO todos (text, completed, user_id) VALUES ($1, $2, $3)
	RETURNING id, created_at
`
	err := r.DB.QueryRowContext(ctx, query, todo.Text, todo.Completed, todo.UserID).Scan(
		&todo.ID,
		&todo.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r repository) GetById(ctx context.Context, id int) (*Todo, error) {
	query := `
	SELECT id, text, completed, user_id, created_at FROM todos WHERE id = $1
`
	var todo Todo
	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&todo.ID,
		&todo.Text,
		&todo.Completed,
		&todo.UserID,
		&todo.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r repository) GetAll(ctx context.Context) ([]Todo, error) {
	query := `
	SELECT id, text, completed, user_id, created_at FROM todos
`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	res := make([]Todo, 0)
	for rows.Next() {
		var todo Todo
		err := rows.Scan(
			&todo.ID,
			&todo.Text,
			&todo.Completed,
			&todo.UserID,
			&todo.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		res = append(res, todo)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return res, nil
}

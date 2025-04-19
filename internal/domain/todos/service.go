package todos

import (
	"context"
	"errors"
	"github.com/lucianboboc/todo-api/internal/intrastructure/database"
)

type Service interface {
	Create(ctx context.Context, todo *Todo) error
	GetById(ctx context.Context, id int) (*Todo, error)
	GetAll(ctx context.Context) ([]Todo, error)
}

func NewService(r Repository) Service {
	return service{
		repository: r,
	}
}

type service struct {
	repository Repository
}

func (s service) Create(ctx context.Context, todo *Todo) error {
	return s.repository.Create(ctx, todo)
}

func (s service) GetById(ctx context.Context, id int) (*Todo, error) {
	todo, err := s.repository.GetById(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrNoRecordsFound):
			return nil, ErrTodoNotFound
		}
	}
	return todo, nil
}

func (s service) GetAll(ctx context.Context) ([]Todo, error) {
	return s.repository.GetAll(ctx)
}

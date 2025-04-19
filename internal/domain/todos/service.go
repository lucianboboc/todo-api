package todos

import "context"

type Service interface {
	Create(ctx context.Context, todo Todo) error
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

func (s service) Create(ctx context.Context, todo Todo) error {
	return nil
}

func (s service) GetById(ctx context.Context, id int) (*Todo, error) {
	return nil, nil
}

func (s service) GetAll(ctx context.Context) ([]Todo, error) {
	return nil, nil
}

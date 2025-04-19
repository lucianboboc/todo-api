package users

import (
	"context"
	"errors"
	"github.com/lucianboboc/todo-api/internal/intrastructure/database"
	"github.com/lucianboboc/todo-api/internal/intrastructure/security"
)

type Service interface {
	Create(ctx context.Context, user *User, password string) error
	GetUsers(ctx context.Context) ([]User, error)
	GetByID(ctx context.Context, id int) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}

func NewService(securityService security.Service, repository Repository) Service {
	return service{
		securityService: securityService,
		repository:      repository,
	}
}

type service struct {
	securityService security.Service
	repository      Repository
}

func (s service) Create(ctx context.Context, user *User, password string) error {
	passwordHash, err := s.securityService.HashPassword(password)
	if err != nil {
		return err
	}
	user.PasswordHash = passwordHash
	err = s.repository.Create(ctx, user)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrRecordAlreadyExists):
			return ErrUserAlreadyExists
		}
		return err
	}
	return nil
}

func (s service) GetUsers(ctx context.Context) ([]User, error) {
	return s.repository.GetUsers(ctx)
}

func (s service) GetByID(ctx context.Context, id int) (*User, error) {
	user, err := s.repository.GetByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrNoRecordsFound):
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (s service) GetByEmail(ctx context.Context, email string) (*User, error) {
	user, err := s.repository.GetByEmail(ctx, email)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrNoRecordsFound):
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

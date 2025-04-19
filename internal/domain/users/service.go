package users

import (
	"context"
	"github.com/lucianboboc/todo-api/internal/intrastructure/security"
)

type Service interface {
	Create(ctx context.Context, user *User, password string) error
	GetByID(ctx context.Context, id int) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
}

func NewService(passwordService security.Service, repository Repository) Service {
	return service{
		passwordService: passwordService,
		repository:      repository,
	}
}

type service struct {
	passwordService security.Service
	repository      Repository
}

func (s service) Create(ctx context.Context, user *User, password string) error {
	passwordHash, err := s.passwordService.HashPassword(password)
	if err != nil {
		return err
	}
	user.PasswordHash = passwordHash
	return s.repository.Create(ctx, user)
}

func (s service) GetByID(ctx context.Context, id int) (*User, error) {
	return s.repository.GetByID(ctx, id)
}

func (s service) GetByEmail(ctx context.Context, email string) (*User, error) {
	return s.repository.GetByEmail(ctx, email)
}

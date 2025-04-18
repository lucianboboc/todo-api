package auth

import (
	"context"
	"time"
	"todo_api/internal/pkg/jsonwebtoken"
	"todo_api/internal/pkg/password"
	"todo_api/internal/services/users"
)

type Service interface {
	Login(ctx context.Context, email, password string) (string, error)
}

func NewService(usersService users.Service, passwordService password.Service, jwtService jsonwebtoken.Service) Service {
	return service{
		usersService:    usersService,
		passwordService: passwordService,
		jwtService:      jwtService,
	}
}

type service struct {
	usersService    users.Service
	passwordService password.Service
	jwtService      jsonwebtoken.Service
}

func (s service) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.usersService.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	err = s.passwordService.ComparePassword(password, user.PasswordHash)
	if err != nil {
		return "", err
	}

	token, err := s.jwtService.GenerateToken(user.ID, 1*time.Hour)
	if err != nil {
		return "", err
	}

	return token, nil
}

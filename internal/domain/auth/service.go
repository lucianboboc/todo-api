package auth

import (
	"context"
	"github.com/lucianboboc/todo-api/internal/domain/users"
	"github.com/lucianboboc/todo-api/internal/intrastructure/jsonwebtoken"
	"github.com/lucianboboc/todo-api/internal/intrastructure/security"
	"time"
)

type Service interface {
	Login(ctx context.Context, email, password string) (string, error)
}

func NewService(usersService users.Service, securityService security.Service, jwtService jsonwebtoken.Service) Service {
	return service{
		usersService:    usersService,
		securityService: securityService,
		jwtService:      jwtService,
	}
}

type service struct {
	usersService    users.Service
	securityService security.Service
	jwtService      jsonwebtoken.Service
}

func (s service) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.usersService.GetByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	err = s.securityService.ComparePassword(password, user.PasswordHash)
	if err != nil {
		return "", err
	}

	token, err := s.jwtService.GenerateToken(user.ID, 1*time.Hour)
	if err != nil {
		return "", err
	}

	return token, nil
}

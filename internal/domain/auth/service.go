package auth

import (
	"context"
	"github.com/lucianboboc/todo-api/internal/domain/users"
	"github.com/lucianboboc/todo-api/internal/intrastructure/jsonwebtoken"
	"github.com/lucianboboc/todo-api/internal/intrastructure/security"
	"log/slog"
	"time"
)

type Service interface {
	Login(ctx context.Context, email, password string) (string, error)
}

func NewService(usersService users.Service, securityService security.Service, jwtService jsonwebtoken.Service, logger *slog.Logger) Service {
	return service{
		usersService:    usersService,
		securityService: securityService,
		jwtService:      jwtService,
		logger:          logger,
	}
}

type service struct {
	usersService    users.Service
	securityService security.Service
	jwtService      jsonwebtoken.Service
	logger          *slog.Logger
}

func (s service) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.usersService.GetByEmail(ctx, email)
	if err != nil {
		s.logger.Error(err.Error())
		return "", ErrUserUnauthorized
	}

	err = s.securityService.ComparePassword(password, user.PasswordHash)
	if err != nil {
		s.logger.Error(err.Error())
		return "", ErrUserUnauthorized
	}

	token, err := s.jwtService.GenerateToken(user.ID, 1*time.Hour)
	if err != nil {
		return "", err
	}

	return token, nil
}

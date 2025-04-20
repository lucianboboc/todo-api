package main

import (
	"github.com/lucianboboc/todo-api/config"
	"github.com/lucianboboc/todo-api/internal/domain/auth"
	"github.com/lucianboboc/todo-api/internal/domain/todos"
	"github.com/lucianboboc/todo-api/internal/domain/users"
	"github.com/lucianboboc/todo-api/internal/intrastructure/database"
	"github.com/lucianboboc/todo-api/internal/intrastructure/jsonwebtoken"
	"github.com/lucianboboc/todo-api/internal/intrastructure/security"
	"github.com/lucianboboc/todo-api/internal/transport/http/handlers/authhandler"
	"github.com/lucianboboc/todo-api/internal/transport/http/handlers/todoshandler"
	"github.com/lucianboboc/todo-api/internal/transport/http/handlers/usershandler"
	"log/slog"
	"net/http"
)

func newServeMux(db database.Database, conf *config.Config, logger *slog.Logger) *http.ServeMux {
	mux := http.NewServeMux()

	usersRepository := users.NewRepository(db)
	todosRepository := todos.NewRepository(db)

	securityService := security.NewService()
	usersService := users.NewService(securityService, usersRepository)
	jwtService := jsonwebtoken.NewService(conf.JWTSecret)
	todosService := todos.NewService(todosRepository)
	authService := auth.NewService(usersService, securityService, jwtService)

	authHandler := authhandler.NewHandler(authService, logger)
	authHandler.RegisterRoutes(mux)

	usersHandler := usershandler.NewHandler(usersService, jwtService, logger)
	usersHandler.RegisterRoutes(mux)

	todosHandler := todoshandler.NewHandler(todosService, usersService, jwtService, logger)
	todosHandler.RegisterRoutes(mux)
	return mux
}

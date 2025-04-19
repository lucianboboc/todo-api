package main

import (
	"github.com/lucianboboc/todo-api/config"
	"github.com/lucianboboc/todo-api/internal/domain/auth"
	"github.com/lucianboboc/todo-api/internal/domain/todos"
	"github.com/lucianboboc/todo-api/internal/domain/users"
	"github.com/lucianboboc/todo-api/internal/intrastructure/database"
	"github.com/lucianboboc/todo-api/internal/intrastructure/jsonwebtoken"
	"github.com/lucianboboc/todo-api/internal/intrastructure/security"
	"github.com/lucianboboc/todo-api/internal/transport/http/handlers"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	conf, err := config.LoadEnv()
	if err != nil {
		panic(err)
	}

	db, err := database.NewDatabase(conf.Database)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	usersRepository := users.NewRepository(db)
	todosRepository := todos.NewRepository(db)

	securityService := security.NewService()
	usersService := users.NewService(securityService, usersRepository)
	jwtService := jsonwebtoken.NewService(conf.JWTSecret)
	todosService := todos.NewService(todosRepository)
	authService := auth.NewService(usersService, securityService, jwtService)

	authHandler := handlers.NewAuthHandler(authService, logger)
	authHandler.RegisterRoutes(mux)

	usersHandler := handlers.NewUserHandler(usersService, jwtService, logger)
	usersHandler.RegisterRoutes(mux)

	todosHandler := handlers.NewTodoHandler(todosService, usersService, jwtService, logger)
	todosHandler.RegisterRoutes(mux)

	app := newApplication(mux, conf, logger)
	app.Start()
}

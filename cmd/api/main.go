package main

import (
	"fmt"
	"github.com/lucianboboc/todo-api/config"
	"github.com/lucianboboc/todo-api/internal/pkg/database"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	conf, err := config.LoadEnv()
	if err != nil {
		panic(err)
	}

	db, err := database.Open(conf.Database)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	s := http.Server{
		Addr:         fmt.Sprintf(":%d", conf.Port),
		Handler:      http.DefaultServeMux, //TODO: Add routes for users, todos and auth
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 20,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("Starting server", slog.String("addr", s.Addr))
	err = s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

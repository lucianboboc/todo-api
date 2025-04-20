package main

import (
	"github.com/lucianboboc/todo-api/config"
	"github.com/lucianboboc/todo-api/internal/intrastructure/database"
	"log/slog"
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

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	mux := newServeMux(db, conf, logger)
	app := newApplication(mux, conf, logger)
	app.Start()
}

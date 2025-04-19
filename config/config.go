package config

import (
	"errors"
	"github.com/joho/godotenv"
	"github.com/lucianboboc/todo-api/internal/intrastructure/database"
	"github.com/lucianboboc/todo-api/internal/intrastructure/env"
)

type Config struct {
	Database  *database.PostgresConfig
	JWTSecret string
	Port      int
}

func LoadEnv() (*Config, error) {
	var cfg Config
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// Parse database config
	dbHost := env.GetString("DB_HOST", "")
	dbPort := env.GetInt("DB_PORT", 0)
	dbUser := env.GetString("DB_USER", "")
	dbPass := env.GetString("DB_PASS", "")
	dbName := env.GetString("DB_NAME", "")
	if dbHost == "" || dbPort == 0 || dbUser == "" || dbPass == "" || dbName == "" {
		return nil, errors.New("missing database .env values")
	}
	cfg.Database = database.NewPostgresConfig(dbHost, dbPort, dbUser, dbPass, dbName)

	// Parse jwt secret
	cfg.JWTSecret = env.GetString("JWT_SECRET", "")

	// Parse port
	cfg.Port = env.GetInt("PORT", 8080)

	return &cfg, nil
}

package database

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

func NewPostgresConfig(host string, port int, username string, password string, dbName string) *PostgresConfig {
	return &PostgresConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		DBName:   dbName,
	}
}

func (c *PostgresConfig) String() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.DBName,
	)
}

func Open(config *PostgresConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", config.String())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

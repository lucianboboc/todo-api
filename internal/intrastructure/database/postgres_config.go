package database

import "fmt"

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

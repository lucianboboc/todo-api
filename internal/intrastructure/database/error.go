package database

import (
	"errors"
	"github.com/jackc/pgx/v5/pgconn"
)

var ErrRecordAlreadyExists = errors.New("record already exists")
var ErrNoRecordsFound = errors.New("no records found")

func IsUniqueConstraint(err error) bool {
	var pgErr *pgconn.PgError
	switch {
	case errors.As(err, &pgErr):
		return pgErr.Code == "23505"
	}
	return false
}

package postgres

import (
	"fmt"
	"os"

	"github.com/io-m/app-hyphen/pkg/constants"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresConnection() (*sqlx.DB, error) {
	connStr := os.Getenv(constants.POSTGRES_CONNECTION)
	// Instead of sql.Open, we use sqlx.Connect. This method combines sql.Open and db.Ping for us.
	db, err := sqlx.Connect(constants.DRIVER, connStr)
	if err != nil {
		return nil, fmt.Errorf("error connecting to Postgres: %w", err)
	}
	return db, nil
}

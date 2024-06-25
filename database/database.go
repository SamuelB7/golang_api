package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func NewPostgreSQLStorage(config pgx.ConnConfig) (*sql.DB, error) {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

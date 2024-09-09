package main

import (
	"database/sql"
	"go-api/cmd/api"
	"go-api/database"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file on main.go")
	}

	db, err := database.NewPostgreSQLStorage(pgx.ConnConfig{
		Config: pgconn.Config{
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Database: os.Getenv("POSTGRES_DB"),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewServer(":3333", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(database *sql.DB) {
	err := database.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected!")
}

package models

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var con *pgx.Conn

func Connect() (*pgx.Conn, error) {
	//DbDriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", DbUser, DbPassword, DbHost, DbPort, DbName)
	db, err := pgx.Connect(context.Background(), connString)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error connecting to database: %v\n", err)
		return nil, err
	}

	con = db
	fmt.Println("Connected to the database")

	return db, nil
}

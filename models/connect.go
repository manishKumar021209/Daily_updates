package models

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var con *pgx.Conn

func Connect() (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), "postgres://postgres:password@localhost:5432/practice")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error in connetcing to database %v\n", err)
		os.Exit(1)
	}
	con = db
	fmt.Println("connected to database")

	return db, nil
}

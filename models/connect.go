package models

import (
	"database/sql"
	"fmt"
	"log"
)

var con *sql.DB

func Connect() (*sql.DB, error) {
	//connstr := "host=localhost port=5432 user=postgres password=password dbname=newdb sslmode=disable"
	// connstr := "postgres://demo1:password@localhost:5432/newdb?sslmode=disable"
	// connstr = "postgres://manish:password@localhost/mydb?sslmode=disable"
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=password dbname=newdb sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database")
	con = db
	return db, nil

}

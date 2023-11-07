package main

import (
	"context"
	"crud/controllers"
	"crud/models"
	"log"
	"net/http"
)

func main() {
	mux := controllers.Register()
	db, err := models.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close(context.Background())

	log.Fatal(http.ListenAndServe(":3000", mux))

}

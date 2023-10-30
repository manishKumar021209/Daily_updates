package main

import (
	"api/controllers"
	"api/models"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	mux := controllers.Register()
	db, err := models.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Serving....")
	log.Fatal(http.ListenAndServe(":3000", mux))
}

package main

import (
	"context"

	"crud/controllers"
	"crud/models"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	mux := controllers.Register()
	db, err := models.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// ApiHost := os.Getenv("API_HOST")
	ApiPort := os.Getenv("API_PORT")

	defer db.Close(context.Background())

	// Provide the full address including the host and port
	address := fmt.Sprintf(":%s", ApiPort)
	log.Fatal(http.ListenAndServe(address, mux))
}

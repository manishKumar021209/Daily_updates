package main

import (
	"auth/handlers"
	"fmt"
	"net/http"
)

func main() {
	mux := handlers.Register()
	fmt.Println("server is running at port 3001")
	http.ListenAndServe(":3001", mux)
}

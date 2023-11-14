package controllers

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()

	// Apply AuthMiddleware for protected routes
	protectedRoutes := http.NewServeMux()

	// Apply TokenAuth only to specific routes
	//protectedRoutes.Handle("/create", TokenAuth(http.HandlerFunc(CreateStudent)))
	protectedRoutes.Handle("/read", TokenAuth(http.HandlerFunc(ReadAll)))
	protectedRoutes.Handle("/delete", TokenAuth(http.HandlerFunc(DeleteByID)))

	// Routes without TokenAuth
	protectedRoutes.HandleFunc("/create", CreateStudent)
	protectedRoutes.HandleFunc("/updatestudent", updateS)
	protectedRoutes.HandleFunc("/updateparent", updateP)
	protectedRoutes.HandleFunc("/createuser", createUser)
	protectedRoutes.HandleFunc("/login", LoginUser)

	// Apply AuthMiddleware to protected routes
	mux.Handle("/", protectedRoutes)

	return mux
}

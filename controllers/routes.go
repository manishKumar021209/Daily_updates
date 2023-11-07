package controllers

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/create", CreateStudent)
	mux.HandleFunc("/read", ReadAll)

	return mux
}

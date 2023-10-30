package controllers

import "net/http"

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", ping)
	mux.HandleFunc("/", create)
	mux.HandleFunc("/get", readAll)
	mux.HandleFunc("/read", readByName)
	mux.HandleFunc("/delete", deletebyName)

	return mux
}

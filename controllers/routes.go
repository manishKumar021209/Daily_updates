package controllers

import (
	"net/http"
)

func Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/create", CreateStudent)
	mux.HandleFunc("/read", ReadAll)
	mux.HandleFunc("/updatestudent", updateS)
	mux.HandleFunc("/updateparent", updateP)
	mux.HandleFunc("/delete", DeleteByID)

	return mux
}

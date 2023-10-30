package controllers

import (
	"api/views"
	"encoding/json"
	"net/http"
)

// func ping() http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		if r.Method == http.MethodGet {
// 			data := views.Response{ //importing the structs file
// 				Code: http.StatusOK,
// 				Body: "pong",
// 			}
// 			json.NewEncoder(w).Encode(data)
// 		}
// 	}

// }

func ping(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data := views.Response{
			Code: http.StatusOK,
			Body: "pong",
		}
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(data)
	}
}

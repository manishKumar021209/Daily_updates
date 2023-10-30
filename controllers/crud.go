package controllers

import (
	//"encoding/json"

	"api/models"
	"api/views"
	"encoding/json"
	"fmt"

	"log"
	"net/http"
)

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		data := views.PostRequest{}
		json.NewDecoder(r.Body).Decode(&data)
		fmt.Println(data)
		if err := models.CreateTodo(data.Name, data.Todo); err != nil {
			w.Write([]byte("Some error"))
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(data)

		// } else if r.Method == http.MethodGet {

		// 	data, err := models.ReadAll()
		// 	if err != nil {
		// 		w.Write([]byte(err.Error()))
		// 	}
		// 	w.WriteHeader(http.StatusOK)
		// 	json.NewEncoder(w).Encode(data)
		// } else if r.Method == http.MethodGet {
		// 	name := r.URL.Query().Get("name")
		// 	data, err := models.ReadByName(name)
		// 	if err != nil {
		// 		w.Write([]byte(err.Error()))
		// 	}
		// 	w.WriteHeader(http.StatusOK)
		// 	json.NewEncoder(w).Encode(data)
		// } else if r.Method == http.MethodDelete {
		// 	name := r.URL.Path[1:]
		// 	if err := models.DeleteTodo(name); err != nil {
		// 		w.Write([]byte("Some error"))
		// 		return
		// 	}
		// 	w.WriteHeader(http.StatusOK)
		// 	json.NewEncoder(w).Encode(struct {
		// 		Status string `json:"status"`
		// 	}{"Item deleted"})
		// }
	}
}

func readAll(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		data, err := models.ReadAll()

		if err != nil {
			w.Write([]byte(err.Error()))
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)

	}
}

func readByName(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		name := r.URL.Query().Get("name")

		data, err := models.ReadByName(name)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte(err.Error()))
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(data)
	}
}

func deletebyName(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {

		name := r.URL.Query().Get("name")

		if err := models.DeleteTodo(name); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Some error"))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct {
			Status string `json:"status"`
		}{"Item deleted"})
	}
}

// func deletebyName(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodDelete {
// 		// Extract the name from the URL path
// 		name := mux.Vars(r)["name"]

// 		if name == "" {
// 			w.WriteHeader(http.StatusBadRequest)
// 			w.Write([]byte("Name is missing in the URL"))
// 			return
// 		}

// 		if err := models.DeleteTodo(name); err != nil {
// 			log.Println(err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			w.Write([]byte("Some error"))
// 			return
// 		}

// 		w.WriteHeader(http.StatusOK)
// 		json.NewEncoder(w).Encode(struct {
// 			Status string `json:"status"`
// 		}{"Item deleted"})
// 	}
// }

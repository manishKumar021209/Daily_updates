package controllers

import (
	"crud/auth"
	"crud/models"
	"crud/views"
	"encoding/json"
	"fmt"

	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var insert struct {
			UserName string `json:"username"`
			Password string `json:"password"`
		}

		var err error
		if err = json.NewDecoder(r.Body).Decode(&insert); err != nil {
			http.Error(w, "Failed to decode request body", http.StatusBadRequest)
			return
		}
		var usr views.User
		usr.UserName = insert.UserName
		bytes, err := bcrypt.GenerateFromPassword([]byte(insert.Password), 14)
		usr.Password = string(bytes)

		err = models.InsertUser(&usr)
		if err != nil {
			if err != nil {
				http.Error(w, "Failed to create user", http.StatusBadRequest)
				return
			}
		}

	}

}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var details struct {
			UserName string `json:"username"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&details); err != nil {
			http.Error(w, "failed to decode request body", http.StatusBadRequest)
			return
		}

		user, err := models.GetUser(details.UserName)
		if err != nil {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}
		// fmt.Println("1")
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(details.Password))
		if err != nil {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}
		//fmt.Println("2")

		// Generate JWT token upon successful login
		token, err := auth.GenerateJWT(details.UserName)
		if err != nil {
			http.Error(w, "Failed to generate JWT token", http.StatusInternalServerError)
			return
		}

		// Return the generated token in the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"token": token})

		fmt.Printf("User with username %s has logged in", details.UserName)
	}
}

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	user, exists := ctx.Value("user").(string)
	if !exists {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %s!", user)
}

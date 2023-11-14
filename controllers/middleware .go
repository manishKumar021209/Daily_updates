package controllers

import (
	"context"
	"crud/auth"
	"fmt"
	"net/http"
	"strings"
)

func TokenAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}
		fmt.Println("Token:", tokenString)
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		username, err := auth.ExtractJWT(tokenString)
		if err != nil {
			if strings.Contains(err.Error(), "Token is expired") {
				http.Error(w, "Token has expired", http.StatusUnauthorized)
			} else {
				fmt.Println("Error extracting JWT:", err)
				http.Error(w, "Invalid token", http.StatusUnauthorized)
			}
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, "user", username)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

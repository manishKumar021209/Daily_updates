package handlers

import "net/http"

func Register() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("/login", Login)
	r.HandleFunc("/welcome", Welcome)
	r.HandleFunc("/logout", Logout)
	return r

}

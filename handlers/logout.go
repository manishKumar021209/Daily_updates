package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("user have logged out1")

	c, err := r.Cookie("session_token")

	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//fmt.Println("user have logged out2")

	sessionToken := c.Value
	//remove session of a user from sessions map to log the user out
	delete(sessions, sessionToken)

	//we also need to let the client know that cookie is expired
	//so we set its value to empty and expiry time to current time

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   "",
		Expires: time.Now(),
	})
	//fmt.Println("user have logged out")
	w.Write([]byte(fmt.Sprint("logged out!!")))

}

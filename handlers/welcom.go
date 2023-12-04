package handlers

import (
	"fmt"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("welcome1")

	//obtain session tokenfrom request cookies which comes with every request
	c, err := r.Cookie("session_token")
	if err != nil {
		if err == http.ErrNoCookie {
			//if the cookie is not set return unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//fmt.Println("welcome2")

	sessionToken := c.Value
	//fmt.Println(sessionToken)

	userSession, exists := sessions[sessionToken]
	//fmt.Println(exists)

	if !exists {
		//if the session token is not present in sessions map
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	//fmt.Println("welcome3")

	// If the session is present, but has expired, we can delete the session, and return
	// an unauthorized status
	if userSession.isExpired() {
		delete(sessions, sessionToken)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	//if the session  is valid return the welcome message
	w.Write([]byte(fmt.Sprintf("welcome %s!!", userSession.username)))
	//fmt.Println("welcome")
}

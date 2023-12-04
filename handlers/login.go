package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// using a map to store user's login info
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// a map to store sessions
var sessions = map[string]session{}

type session struct {
	username string
	expiry   time.Time
}

func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("user logged in1")

	var creds Credentials

	//get the json body and decode into type Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		//if the structure of the provided body is wrong then return err
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//fmt.Println("user logged in2")
	//get the expected password from users map
	expectedPassword, ok := users[creds.Username]
	//fmt.Println("user logged in3")

	if !ok || expectedPassword != creds.Password {
		//fmt.Println("user logged in4")

		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(fmt.Sprintf("invalid user %s!!", creds.Username)))
		return
	}
	//fmt.Println("user logged in5")
	//create a new random session if the user login
	sessionToken := uuid.NewString()
	//fmt.Println(sessionToken)
	expiresAt := time.Now().Add(120 * time.Second)

	//store the generated token in the session map
	sessions[sessionToken] = session{
		username: creds.Username,
		expiry:   expiresAt,
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session_token",
		Value:   sessionToken,
		Expires: expiresAt,
	})
	//fmt.Println("user logged in")
	w.Write([]byte(fmt.Sprintf("logged in %s!!", creds.Username)))
}

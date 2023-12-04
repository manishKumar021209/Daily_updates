package main

import (
	"fmt"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	cookie := http.Cookie{}
	cookie.Name = "accessToken"
	cookie.Value = "ek so pichhattar"
	cookie.Expires = time.Now().Add(2 * time.Minute)
	cookie.Secure = false
	cookie.HttpOnly = true
	cookie.Path = "/"
	http.SetCookie(w, &cookie)

	cookie2 := http.Cookie{}
	cookie2.Name = "page"
	cookie2.Value = "golang"
	cookie2.Expires = time.Now().Add(2 * time.Minute)
	cookie2.Secure = false
	cookie2.HttpOnly = true
	cookie2.Path = "/"
	http.SetCookie(w, &cookie2)

	fmt.Fprintf(w, "this is cookie!!\n")

}

func printCookie(w http.ResponseWriter, r *http.Request) {
	var retunStr string
	for _, cookie := range r.Cookies() {
		retunStr = retunStr + cookie.Name + ":" + cookie.Value + "\n"
	}
	fmt.Fprintf(w, retunStr)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/print", printCookie)

	fmt.Println("server is running at port 3000")
	http.ListenAndServe(":3000", nil)
}

package main

import (
	"fmt"
	"net/http"
	"github.com/satori/go.uuid"
)

func main() {
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session-id",
			Value: id.String(),
			Secure: true, // https only
			HttpOnly: true, // cannot be accesed by js
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "favicon.ico")
}

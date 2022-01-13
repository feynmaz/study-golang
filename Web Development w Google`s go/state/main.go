package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/", index)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {

	//http.SetCookie(w, &http.Cookie{
	//	Name:  "my-cookie",
	//	Value: "0",
	//})

	cookie, err := r.Cookie("my-cookie")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "my-cookie",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
	count++
	cookie.Value = strconv.Itoa(count)

	http.SetCookie(w, cookie)

	_, err = io.WriteString(w, cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "favicon.ico")
}

package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/dog.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	io.WriteString(w, `
	<img src="/dog.jpg">
	`)
}

func dogPic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("dog.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}

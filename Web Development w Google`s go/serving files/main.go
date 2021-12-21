package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")

	io.WriteString(w, `
	<!--not serving-->
	<img src="https://www.collinsdictionary.com/images/thumb/dog_230497594_250.jpg?version=4.0.207">
	`)
}

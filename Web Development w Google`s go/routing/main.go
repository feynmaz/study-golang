package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "dog dog doggy")
}

type hotcat int

func (c hotcat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "cat cat kitty")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()

	// /dog/anything/else
	mux.Handle("/dog/", d)

	// /cat
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}

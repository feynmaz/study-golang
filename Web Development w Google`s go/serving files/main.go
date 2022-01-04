package main

import (
	"html/template"
	"io"
	"net/http"
)


var tpl *template.Template

func init () {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)

	http.Handle("/media/", http.StripPrefix("/media", http.FileServer(http.Dir("media/"))))
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "dog.html", nil)
}
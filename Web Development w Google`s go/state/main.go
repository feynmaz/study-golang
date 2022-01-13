package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

type Person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.HandleFunc("/favicon.ico", faviconHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at index: ", r.Method)
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at bar: ", r.Method)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusMovedPermanently)
}

func barred(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at barred: ", r.Method)
	err := tpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "favicon.ico")
}

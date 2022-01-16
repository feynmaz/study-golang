package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/satori/go.uuid"
	"html/template"
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

type user struct {
	UserName string
	FirstName string
	LastName string
}

var tpl *template.Template
var dbUsers = make(map[string]user) // user ID - user
var dbSessions = make(map[string]string) // session ID - user ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}


func main() {
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		sID := uuid.NewV4()
		c = &http.Cookie{
			Name: "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}

	var u user
	if userName, ok := dbSessions[c.Value]; ok {
		u = dbUsers[userName]
	}
	
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Fatalf("Cannot parse form: %v", err)
		}
		err = decoder.Decode(&u, r.Form)
		if err != nil {
			log.Fatalf("Cannot decode form: %v", err)
		}
		dbSessions[c.Value] = u.UserName
		dbUsers[u.UserName] = u
	}
	tpl.ExecuteTemplate(w, "index.html", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} 
	un, ok := dbSessions[c.Value]
	if !ok {
		fmt.Println("not ok")
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.html", u)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "favicon.ico")
}

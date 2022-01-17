package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	uuid "github.com/satori/go.uuid"
)

var decoder = schema.NewDecoder()

type user struct {
	UserName  string
	Password string
	FirstName string
	LastName  string
}

var tpl *template.Template
var dbSessions = make(map[string]string) // session id - user id
var dbUsers = make(map[string]user) // user id - user 

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(r)
	tpl.ExecuteTemplate(w, "index.html", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	u := getUser(r)
	if !alreadyLoggedIn(r){
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "bar.html", u)
}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		
		// get form values
		var u user
		err := r.ParseForm()
		if err != nil {
			log.Fatalf("Cannot parse form: %v", err)
		}
		err = decoder.Decode(&u, r.Form)
		if err != nil {
			log.Fatalf("Cannot decode form: %v", err)
		}

		// is username taken
		if _, ok := dbUsers[u.UserName]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name: "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = u.UserName

		// save user
		dbUsers[u.UserName] = u

		// redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.html", nil)
}


func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "favicon.ico")
}
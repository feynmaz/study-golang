package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserName  string
	Password []byte
	FirstName string
	LastName  string
	Role string
}

type session struct{
	un string
	lastActivity time.Time
}

var tpl *template.Template
var dbUsers = make(map[string]user) // user id - user 
var dbSessions = make(map[string]session) // session id - session
var dbSessionsCleaned time.Time
const sessionLength int = 30


func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
	bs, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	dbUsers["test@test.com"] = user{"test@test.com", bs, "James", "Bond", "007"}
}

func main() {
	http.HandleFunc("/favicon.ico", faviconHandler)
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/login", login)
	http.HandleFunc("logout", logout)
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
	if u.Role != "007" {
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	tpl.ExecuteTemplate(w, "bar.html", u)
}

func login(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var u user

	if r.Method == http.MethodPost{
		un := r.FormValue("UserName")
		p := r.FormValue("password")

		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		sID := uuid.NewV4()
		c := &http.Cookie{
			Name: "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.html", u)
}

func logout(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r){
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	c, _ := r.Cookie("session")

	// delete session
	delete(dbSessions, c.Value)

	// remove the cookie
	c = &http.Cookie{
		Name: "session",
		Value: "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	if time.Now().Sub(dbSessionsCleaned) > (time.Second * 30) {
		go cleanSessions()
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	var u user

	if req.Method == http.MethodPost {

		un := req.FormValue("UserName")
		p := req.FormValue("Password")
		f := req.FormValue("FirstName")
		l := req.FormValue("LastName")
		r := req.FormValue("role")

		// is username taken
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// create session
		sID := uuid.NewV4()
		c := &http.Cookie{
			Name: "session",
			Value: sID.String(),
		}
		c.MaxAge = sessionLength
		http.SetCookie(w, c)
		dbSessions[c.Value] = session{un, time.Now()}

		// save user
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		u = user{un, bs, f, l, r}
		dbUsers[un] = u

		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
		
	}

	tpl.ExecuteTemplate(w, "signup.html", nil)
}

func cleanSessions() {
	fmt.Println("BEFORE CLEAN")
	showSessions()
	
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) || v.un == "" {
			delete(dbSessions, k)
		}
	}
	dbSessionsCleaned = time.Now()
	fmt.Println("AFTER CLEAN")
	showSessions()
}

func showSessions() {
	fmt.Println("*********")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println("")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "favicon.ico")
}
package main

import (
	"net/http"
)

func getUser(r *http.Request) user {
	var u user

	// get cookie
	c, err := r.Cookie("session")
	if err != nil {
		return u
	}

	// if the user exists already, get user
	if s, ok := dbSessions[c.Value]; ok {
		u = dbUsers[s.un]
	}
	return u
}

func alreadyLoggedIn(r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	s := dbSessions[c.Value]
	_, ok := dbUsers[s.un]
	return ok
}
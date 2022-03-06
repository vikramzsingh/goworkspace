package main

import (
	"net/http"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	var u user

	// check cookie is available ?
	c, err := req.Cookie("session")
	if err != nil {
		return u // empty user
	}

	// if user already exist, get user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func allreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}

	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
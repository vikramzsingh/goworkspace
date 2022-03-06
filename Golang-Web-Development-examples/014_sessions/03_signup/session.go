package main

import "net/http"

func getUser(req *http.Request) user {
	var u user

	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		return u
	}

	// if the user exists already, get user
	if un, ok := dbSession[c.Value]; ok {
		u = dbUser[un]
	}
	return u
}

func alreadyLoggedIn(req *http.Request)  bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}

	un := dbSession[c.Value] // fetching username from map, as Key_Type is setted in cookie Value
	_, ok := dbUser[un]
	return ok // true or false
}
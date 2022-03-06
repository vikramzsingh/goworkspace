package main

import (
	"fmt"
	"net/http"

	"github.com/satori/go.uuid"

)

func main() {
	/*
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	// or
	*/
	mux := http.NewServeMux()
	mux.HandleFunc("/", foo)
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", mux)
}

func foo(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session-id")
	if err != nil {
		id, _ := uuid.NewV4() // UUID --> Universal unique id
		/* or handle error
		if err != nil {
			fmt.Println(err)
		}
		*/
		cookie = &http.Cookie{
			Name:  "session-id",
			Value: id.String(), // converting UUID into string // This string method built-in inside satori package 
			// Secure: true, // leave on when using TLS(https) // only allow cookie to set when using https 
			HttpOnly: true, // cookie is not accessible with JavaScript, it can only accessible with http // provides more security
		}
		http.SetCookie(w, cookie)
	}
	fmt.Println(cookie)
}

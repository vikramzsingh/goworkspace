package main

import (
	"fmt"
	"net/http"

)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, `<h1><a href="/set">set a cookie</a></h1>`)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: "some value",
	})

	fmt.Fprintf(w, `<h1><a href="/read">read cookie</a></h1>`)
}

func read(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		//http.Redirect(w, req, "/set", http.StatusSeeOther) // 303
		return
	}

	fmt.Fprintf(w, `<h1>YOUR COOKIE:<br> %v<h1><h1><a href="/expire">expire cookie</a></h1>`, c)
}

func expire(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther) // 303
		return
	}
	c.MaxAge = -1 // delete cookie 
	http.SetCookie(w, c)
	http.Redirect(w, req, "/read", http.StatusSeeOther) // 303
}
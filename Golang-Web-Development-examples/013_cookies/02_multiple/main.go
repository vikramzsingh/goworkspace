package main

import (
	"fmt"
	"net/http"
	"time"

)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/abundance", abundance)
	http.HandleFunc("/favicon.ico", http.NotFoundHandler().ServeHTTP)
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "my-cookie",
		Value:   "some value",
		Path:    "/", // OPTIONAL
		Expires: time.Now().Add(10 * time.Minute),
	})

	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	c1, err := req.Cookie("my-cookie") // give me cookie with that Name "my-cookie"
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // 400
		//log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #1: ", c1)
	}

	c2, err := req.Cookie("general") // give me cookie with that Name "general"
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // 400
		//log.Println(err)
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #2: ", c2)
	}

	c3, err := req.Cookie("specific") // give me cookie with that Name "specific"
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest) // 400
		//log.Println(err) // for printing on bash
	} else {
		fmt.Fprintln(w, "YOUR COOKIE #1: ", c3)
	}
}

func abundance(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "some other value about general things",
		Path:  "/", // OPTIONAL
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "some other value about specific things",
	})

	fmt.Fprintln(w, "MULTIPLE COOKIES WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

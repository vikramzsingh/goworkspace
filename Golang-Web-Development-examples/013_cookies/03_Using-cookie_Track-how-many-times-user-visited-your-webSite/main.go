package main

import (
	"io"
	"log"
	"net/http"
	"strconv"

)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("my-cookie")

	if err == http.ErrNoCookie {
		cookie = &http.Cookie{ // just created a cookie struct, and stored into above cookie variable
			Name:  "my-cookie",
			Value: "0",
			Path:  "/",
		}
	}

	count, err := strconv.Atoi(cookie.Value) // string-convert Asci-to-int, now count has Int-type VALUE
	if err != nil {
		log.Fatalln(err)
	}
	count++ // Int

	cookie.Value = strconv.Itoa(count) // string-convert Int-to-asci, Now count has string-type Value, stored in Cookie Struct

	http.SetCookie(w, cookie)

	io.WriteString(w, cookie.Value)
}
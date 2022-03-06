package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler()) // always handle favicon.ico
	http.ListenAndServe(":8080", nil)                   // DefaultServeMux is used
}

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	//v := req.FormValue("fname")
	io.WriteString(w, "Do my search: "+v)
}

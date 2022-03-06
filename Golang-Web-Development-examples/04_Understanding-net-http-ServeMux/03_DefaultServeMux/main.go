package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy doggy doggy")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}
func main() {
	var d hotdog
	var c hotcat

	http.Handle("/dog/", d) // type-Handler is d required
	http.Handle("/cat", c)

	http.ListenAndServe(":8080", nil) // when using DefaultServeMux Type-Handler is nil.\
	// http.Handle() implements Handler INTERFACE
	// which leads to run func ServeHTTP()
}

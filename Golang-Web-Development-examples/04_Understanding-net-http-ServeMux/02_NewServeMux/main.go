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
	// var c hotdog --> Also run dogyy doggy code. As type hotdog is attcahed to doggy doggy code
	var c hotcat

	// NewServeMux() return *ServeMux
	// since ServeHTTP(res http.ResponseWriter, req *http.Request) is also attached to *ServeMux
	// Anybody who has this ServeHTTP(res http.ResponseWriter, req *http.Request) method is also of type Handler INTERFACE
	// So, mux is also of Handler type
	// Hence we passing it into ListenAndServe()
	mux := http.NewServeMux()
	mux.Handle("/dog/", d) // "/dog/" means anything comes after the /something/here/sadds can be displayed in browser
	mux.Handle("/cat", c)  // Require Handler // // "/cat" only this cat url can display content
	// mux.HandleFunc("/cat", http.HandlerFunc(c.ServeHTTP))
	// mux.HandleFunc("/cat", c.ServeHTTP)

	http.ListenAndServe(":8080", mux) // As NeWServeMux return *ServeMux, and ServeHTTP(res, req) is attached
	// to this *ServeMux. Hence NewServeMux implicitly implementing Handler INTERFACE

}

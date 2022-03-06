package main

import (
	"fmt"
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "doggy doggy doggy")
}

func c(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "catty catty catty")
}

func main() {
	// type HandlerFunc func(ResponseWriter, *Request)
	// HandlerFunc is like user defined type, whose underlying type is func(ResponseWriter, *Request)
	// Note :- func(ResponseWriter, *Request) is own type
	// Here we are converting d of type func(ResponseWriter, *Request) into HandlerFunc type
	//
	// Note:- func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
	// METHOD ServeHTTP(w ResponseWriter, r *Request) is attached to HandlerFunc
	// since ServeHTTP(w ResponseWriter, r *Request) is a method inside Handler INTERFACE
	// so, Any body who has this ServeHTTP(w ResponseWriter, r *Request) method is also of Handler type
	// Finally d becomes Handler type
	http.Handle("/dog/", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil)
}

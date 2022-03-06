package main

import (
	"fmt"
	"net/http"
)

/*
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
*/

type hotdog int

// any value of hotdog type this function can use in it.
// this func is also of type hotdog type and any value of hotdog type can use thid func.
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
	fmt.Printf("Type of m(m or d are same type) VARIABLE in main func: %T\n", m) // checking type
	fmt.Printf("Type of serveHTTP in ServeHTPP func: %T\n", m.ServeHTTP)         // checking type
}

func main() {
	var d hotdog // This variable is also of type Handler.
	// As METHOD ServeHTTP(ResponseWriter, *Request) is also of type Handler as this METHOD is available in Handler INTERFACE
	// And METHOD ServeHTTP(ResponseWriter, *Request) is also of type hotdog, as attaced to hotdog type.
	fmt.Printf("Type of d VARIABLE in main func: %T\n", d)          // checking type
	fmt.Printf("Type of serveHTTP in main func: %T\n", d.ServeHTTP) // checking type

	// func ListenAndServe(address string, handler Handler) error
	// return error
	// It takes data of type hotdog/Handler. we give it a Handler.
	// Handler INTERFACE has one METHOD ServeHTTP(w http.ResponseWriter, r *http.Request)
	// It just run That METHOD ServeHTTP(w http.ResponseWriter, r *http.Request) for every request that comes into server.
	http.ListenAndServe(":8080", d) // passing VARIABLE d as Handler type
}

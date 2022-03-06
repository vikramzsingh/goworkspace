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

// Any type who has this METHOD SIGNATURE ServeHTTP(ResponseWriter, *Request) is also of type Handler
// And it also Implementing handler INTERFACE

type hotdog int

// METHOD ServeHTTP(ResponseWriter, *Request) is attached to type hotdog.
// This means METHOD ServeHTTP(ResponseWriter, *Request) can use any value of type hotdog.
// This METHOD ServeHTTP(ResponseWriter, *Request) functionality is also availabe to use for any value of type hotdog.

// Any value of type hotdog is also of type Handler. because this METHOD ServeHTTP(ResponseWriter, *Request) is available in Handler INTERFACE
// And any value of serveHttp is also of type hotdog and vice-versa
func (m hotdog) serveHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Any code you want in this func", m)
}

func main() {

}

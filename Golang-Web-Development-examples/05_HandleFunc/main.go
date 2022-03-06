// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
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
	// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	// Its Parameters Require Url/Path and a function of type - func(ResponseWriter, *Request)
	// d and c are two function we created with type of - func(ResponseWriter, *Request)
	http.HandleFunc("/dog/", d) // func(ResponseWriter, *Request) = d(res http.ResponseWriter, req *http.Request)
	http.HandleFunc("/cat", c)  // func(ResponseWriter, *Request) = c(res http.ResponseWriter, req *http.Request)

	http.ListenAndServe(":8080", nil) // DefaultServeMux used because nil is used

}

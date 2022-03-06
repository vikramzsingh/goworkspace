package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil) // nil means DefaultServeMux is used
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Index page")
}

func d(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "doggy doggy doggy")
}

func me(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Vikram vikram vikram")
}

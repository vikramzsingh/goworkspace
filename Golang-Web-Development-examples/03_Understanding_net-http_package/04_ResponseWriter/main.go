package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Header is a type-map :- map[string][]string
	// Header has also these Methods
	w.Header().Set("Vikram-Key", "This is from Vikram")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Add("Vikram-key-1", "This is from Vikram Second Key")
	keyValue := w.Header().Get("Vikram-Key")
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>")
	fmt.Fprintf(w, "<h1>key(Vikram-Key) Value retrived from Header: %v</h1>", keyValue)
}

func main() {
	var d hotdog

	http.ListenAndServe(":8080", d)
}

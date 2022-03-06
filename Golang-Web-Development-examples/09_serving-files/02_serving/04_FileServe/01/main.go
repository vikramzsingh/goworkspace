package main

import (
	"io"
	"net/http"
)

func main() {
	// . inside Dir means current root directory/ project folder
	http.Handle("/", http.FileServer(http.Dir(".")))
	// http.Handle("/", http.FileServer(http.Dir("./images")))
	http.HandleFunc("/dog", dog)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="toby.jpg"> `)
}

package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// when /resources/toby.jpg occur, the closest path is /resources/
	// So /resources/ run
	// It Strips /resources from img-tag, and ask for /toby.jpg
	// The it looks for toby.jpg in ./assets folder or directory
	io.WriteString(w, `<img src="/resources/toby.jpg">`) // we stripped of this /resources, and now we asking for /toby.jpg
}

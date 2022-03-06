package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogpic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/toby.jpg">`) // when /toby.jpg occur then dogpic runs
	// asking for toby.jpg
}

func dogpic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		// http.Error(w, "file not found", http.StatusNotFound)
		return
	}
	defer f.Close()

	fi, err := f.Stat() //return FileInfo, error
	if err != nil {
		http.Error(w, "file not found", 404)
	}

	// ServeContent(res, *req, file-name, file last-time modified, actual-file)
	http.ServeContent(w, req, f.Name(), fi.ModTime(), f) // serves image to browser
}

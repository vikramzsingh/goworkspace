package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		URL         *url.URL
		Submissions map[string][]string // or url.Values
		Header      http.Header
	}{
		Method:      req.Method,
		URL:         req.URL,
		Submissions: req.Form,
		Header:      req.Header,
	}
	// w writes response to web browser
	tpl.ExecuteTemplate(w, "index.html", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.html"))
}

func main() {
	var d hotdog

	http.ListenAndServe(":8080", d)
}

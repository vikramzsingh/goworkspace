package main

import (
	"html/template"
	"log"
	"net/http"

)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	// body
	bs := make([]byte, req.ContentLength)
	//fmt.Println(bs)
	//r := req.Body
	//fmt.Println(r)
	req.Body.Read(bs) // I say request body read into that bite of slice, it read all the request body into that bite
	s := string(bs)
	//fmt.Println(s)

	err := tpl.ExecuteTemplate(w, "index.gohtml", s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}

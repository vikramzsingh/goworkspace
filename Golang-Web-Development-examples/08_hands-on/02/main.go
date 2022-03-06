package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "tpl.gohtml", "Index page")
	if err != nil {
		log.Fatalln(err)
	}
}

func d(w http.ResponseWriter, req *http.Request) {
	s := "doggy doggy doggy"
	err := tpl.ExecuteTemplate(w, "tpl.gohtml", s)
	if err != nil {
		log.Fatalln(err)
	}
}

func me(w http.ResponseWriter, req *http.Request) {
	s := `Vikram vikram vikram`
	err := tpl.ExecuteTemplate(w, "tpl.gohtml", s)
	if err != nil {
		log.Fatalln(err)
	}
}

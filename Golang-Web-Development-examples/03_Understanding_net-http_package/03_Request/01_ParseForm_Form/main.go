package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Before using Form/PostForm, we must call ParseForm()
	// ParseFrom() parse our Form and make data available to us.
	// These METHODS Form/PostForm are only available after ParseForm() is called.
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	// Form field contains the parsed form data, including both URL.
	// PostForm() contains data from Form body only. (/Post data)
	// ExecuteTemplate func fetch data written in the file "index.gohtml" and run
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
	// req.Form and req.PostForm fields return us url.Values
	// type of Values is map
	// type Values map[string][]string --> it is PRE-DEFINED
	// implicitly field req.Form/req.PostForm returns us map of type map[string][]string

	fmt.Println("The form data retrived from html-from: ", req.Form)
	fmt.Println("The fname(Variable retrived from form):  ", req.Form.Get("fname"))
}

var tpl *template.Template

// On loading program it runs once automatically.
func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog // also of Handler type implicitly through METHOD ServeHTTP

	// This method implements Handler INTERFACE as one PARAMETER is of Handler type
	// func ListenAndServe(address string, handler Handler) error
	// Handler INTERFACE run, it has only one method ServeHTTP(w http.ResponseWriter, r *http.Request)
	// So This METHOD ServeHTTP(w http.ResponseWriter, r *http.Request) runs
	http.ListenAndServe(":8080", d)
}

package main

import (
	"fmt"
	"net/http"
	"text/template"

)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/bar", http.HandlerFunc(bar))
	http.HandleFunc("/barred", barred)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("your request method at bar: ", req.Method)
	http.Redirect(w, req, "/", http.StatusSeeOther) // 303 code or After this Method changes to "GET" in foo
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("your request method at barred: ", req.Method)
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound) // 404 code or
		// http.Error(w, err.Error(), 404) 
	}
}

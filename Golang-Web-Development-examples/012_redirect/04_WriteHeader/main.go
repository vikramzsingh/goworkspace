package main

import (
	"fmt"
	"html/template"
	"net/http"

)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar:", req.Method)
	// process form data
	// Explicitly setting location "/", In order to go foo 
 	// WriteHeader 303 see-other, the go to foo and Method changes to GET
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}


func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred: ", req.Method)
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound) // 404
	}
}

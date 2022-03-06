package main

import (
	"fmt"
	"html/template"
	"net/http"

)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	//http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Print("Your request method at foo: ", req.Method, "\n\n")
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar: ", req.Method)
	// process form data
	http.Redirect(w, req, "/", http.StatusMovedPermanently) // 301 code or After this you will permanently move to foo "/", and cannot go to bar again, even if you try to go it only go to foo (YOU CAN check in bash)
	// http.Redirect(w, req, "/", 301) 						// req Method is POST it Will become GET to foo							
}

/*
func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred:", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
*/

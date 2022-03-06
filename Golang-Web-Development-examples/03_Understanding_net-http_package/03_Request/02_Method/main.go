package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// From/PostForm fields are available after this ParseForm()
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		Submissions url.Values
	}{
		req.Method, // these are built-in methods and fields. // asking Request give me METHOD
		req.Form,   // Form field returns url.Values --> map[string][strings]
	}
	tpl.ExecuteTemplate(w, "index.gohtml", data)
	fmt.Println(data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d hotdog

	http.ListenAndServe(":8080", d)
}

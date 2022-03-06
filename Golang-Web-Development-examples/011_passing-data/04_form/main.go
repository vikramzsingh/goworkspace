package main

import (
	"html/template"
	"net/http"
)

type person struct {
	FIRSTNAME  string
	LASTNAME   string
	Subscribed bool //string use to run without == "on"
}

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
	//w.Header().Set("Content-Type", "text/html; charset=utf-8")

	f := req.FormValue("first")
	l := req.FormValue("last")
	// if subscribe is checked,
	// then it return on,
	//And on == on, gives true
	// so, s = true (bool-type)
	s := req.FormValue("subscribe") == "on"
	/*
		data := req.Form.Get("first")
		fmt.Fprintln(w, data)
	*/
	err := tpl.ExecuteTemplate(w, "index.gohtml", person{f, l, s}) // sending struc data to index.gohtml template
	if err != nil {
		http.Error(w, err.Error(), 500)
		// or --> http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

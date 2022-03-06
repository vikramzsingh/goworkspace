package main

import (
	"html/template"
	"net/http"

	"github.com/satori/go.uuid"

)

type user struct{
	UserName string
	Password string
	First string
	Last string 
}

var tpl *template.Template
var dbSession = map[string]string{}
var dbUser = map[string]user{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(req)

	if !alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther) // 303
	}
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {

		// get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		// username taken or not ?
		if _, ok := dbUser[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{ // for accessing in whole block of form subbmission
			Name: "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSession[c.Value] = un

		// store user in dbUser
		u := user{un, p, f, l}
		dbUser[un] = u

		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}// end of form submission
	
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}
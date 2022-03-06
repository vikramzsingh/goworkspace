package main

import (
	"fmt"
	"html/template"
	"net/http"

	uuid "github.com/satori/go.uuid"

)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template

// empty map
var dbUsers = map[string]user{}      // user ID, user
var dbSessions = map[string]string{} // session ID, user ID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		sID, _ := uuid.NewV4() // creating UUID --> Universal unique identity
		c = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}

	// if the user exists already, get user
	// un stores Value of session which is UUID or c.Value
	var u user
	if un, ok := dbSessions[c.Value]; ok { // at this point map is empty because session is mot setted in map dbSessions, comma ok idiom will be false
		u = dbUsers[un] // dbUsers map Key_Type is string and Value_Type is user struct{}, u holds structure of user
	}

	// prossess form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u := user{un, f, l}
		dbSessions[c.Value] = un // setting cookie-Value as Key_Type and username as Value_Type
		dbUsers[un] = u          // setting username as Key_Type and user struct{} as Value_Type
	}

	err = tpl.ExecuteTemplate(w, "index.gohtml", u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound) // 404
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	fmt.Println(c) // for checking if map not setted, then on clicking go to bar redirected to foo
	
	un, ok := dbSessions[c.Value] // getting value from map where we setted cookie-value as Key_Type and username as Value_Type
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	u := dbUsers[un] // retriving struct from map, related to this user 
	err = tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

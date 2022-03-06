package main

import (
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"github.com/satori/go.uuid"
)

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
}

var tpl *template.Template
var dbSessions = map[string]string{}
var dbUsers = map[string]user{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	bs, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	dbUsers["vikram@gmail.com"] = user{"vikram@gmail.com", bs, "vikram", "singh"}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	err := tpl.ExecuteTemplate(w, "index.gohtml", u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound) // 404
		return
	}
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	if !allreadyLoggedIn(req) { // not login send to index page
		http.Redirect(w, req, "/", http.StatusSeeOther) //404
		return
	}
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func signup(w http.ResponseWriter, req *http.Request) {
	if allreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther) // 303
		return
	}

	var u user
	// process form submissions
	if req.Method == http.MethodPost{

		// form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		// username taken ?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden) //403
			return
		}

		// create session
		c, err := req.Cookie("session")
		if err != nil {
			sID, _ := uuid.NewV4()
			c = &http.Cookie{
				Name: "session",
				Value: sID.String(),
			}
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un

		// store user in dbUser
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		u = user{un, bs, f, l}
		dbUsers[un] = u

		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther) // 303
		return
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

func login(w http.ResponseWriter, req *http.Request) {
	if allreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther) // 303
		return
	}

	// process form submission
	if req.Method == http.MethodPost {

		// form values
		un := req.FormValue("username")
		p := req.FormValue("password")

		// is there a user
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden) // 403
			return
		}

		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden) // 403
			return
		}

		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name: "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSessions[c.Value] = un

		http.Redirect(w, req, "/", http.StatusSeeOther) // 303
		return
	}

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

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
var dbUsers = map[string]user{} // or = map[string]user{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
	bs, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	dbUsers["test@test.com"] = user{"test@test.com", bs, "James", "Bond"}
}

func main() {
	/*
	fmt.Println(dbUsers)
	fmt.Printf("Type of dbUsers:  %T\n", dbUsers)
	fmt.Println(dbSessions)
	fmt.Printf("Type of dbSessions:  %T", dbSessions)
	*/

	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/favicon.ico", http.NotFoundHandler().ServeHTTP)
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
	if !allreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther) // 303
	}

	err := tpl.ExecuteTemplate(w, "bar.gohtml", u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound) // 404
	}
}

func signup(w http.ResponseWriter, req *http.Request) {
	if allreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther) //303
		return
	}

	var u user
	// process form submission
	if req.Method == http.MethodPost {

		// process form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		//username taken ?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden) //403
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

		// store user in dbUser map
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError) //500
			return
		}
		u = user{un, bs, f, l}
		dbUsers[un] = u

		//redirect
		http.Redirect(w, req, "/", http.StatusSeeOther) // 303
		return
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func login(w http.ResponseWriter, req *http.Request) {
	if allreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther) // 303
	}

	var u user
	// process form submission
	if req.Method == http.MethodPost {

		// process form values
		un := req.FormValue("username")
		p := req.FormValue("password")

		// is there a username ?
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
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
	tpl.ExecuteTemplate(w, "login.gohtml", u)
}

func logout(w http.ResponseWriter, req *http.Request) {
	if !allreadyLoggedIn(req) { // not loggned
		http.Redirect(w, req, "/", http.StatusSeeOther) //303
		return
	}

	c, _ := req.Cookie("session")
	// delete session // form map
	delete(dbSessions, c.Value)

	// delete cookie
	c = &http.Cookie{
		Name: "session",
		Value: "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	http.Redirect(w, req, "/", http.StatusSeeOther)
}
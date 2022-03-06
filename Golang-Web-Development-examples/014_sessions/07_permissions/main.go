package main

import (
	"html/template"
	"net/http"

	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)


type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
	Role     string
}

var tpl *template.Template
var dbSessions = map[string]string{}
var dbUsers = map[string]user{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	fmt.Println(u)
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	if !allreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther) // 303
		return // go back don't execute beyond it
	}

	if u.Role != "007" { // only 007 can enter in this block
		http.Error(w, "You must be 007 to enter the bar", http.StatusForbidden)
		return
	}
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
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
		r := req.FormValue("role")

		// user taken ?
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

		// store user in dbUsers
		bs, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		u = user{un, bs, f, l, r}
		dbUsers[un] = u

		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther) // 303
		return // go back dont execute further
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", u)
}

func login(w http.ResponseWriter, req *http.Request) {
	if allreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther) //303
		return
	}

	var u user
	// process form submission
	if req.Method == http.MethodPost {

		// process formm values
		un := req.FormValue("username")
		p := req.FormValue("password")

		// is there username ?
		u, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password", http.StatusForbidden) //403
			return
		}

		// does the entered password match the stored password?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password", http.StatusForbidden) //403
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

		//redirect
		http.Redirect(w, req, "/", http.StatusSeeOther) //303
		return
	}
	tpl.ExecuteTemplate(w, "login.gohtml", u)
}

func logout(w http.ResponseWriter, req *http.Request) {
	st := allreadyLoggedIn(req)
	fmt.Println(st)

	if !allreadyLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther) //303
		return
	}

	fmt.Println("inside logout")
	// get cookie
	c, _ := req.Cookie("session")

	// delete session
	delete(dbSessions, c.Value)

	//remove cookiie
	c = &http.Cookie{
		Name: "session",
		Value: "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)

	// redirect
	http.Redirect(w, req, "/", http.StatusSeeOther) //303
}
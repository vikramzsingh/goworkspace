package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"Golang-Web-Development/020_mondodb/03_post-delete/models"
	"encoding/json"
	"log"
	"fmt"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	// added router
	r.POST("/user", createUser)
	//  added route plus parameter
	r.DELETE("/user/:id", deleteUser) // default is GET
	http.ListenAndServe(":8080", r)
}


// juliansmidth require three parameter
func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="UTF-8">
<title>Index</title>
</head>
<body>
<a href="/user/9872309847">GO TO: http://localhost:8080/user/9872309847</a>
</body>
</html>
	`
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

// getUser also recieve id from url of the function /user/:id
func getUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	u := models.User{
		Name: "James",
		Gender: "Male",
		Age: 54,
		Id: p.ByName("id"), // when link at index is triggered, /user/:id url is its target with id (9872309847), and getUser func is executed
	}

	// Marshal into JSON
	uj, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	w.Write(uj)
	// fmt.Fprintf(w, "%s\n", uj)
}

func createUser(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	u := models.User{} // needed variable to store JSON data in GO

	// get JSON data from request body
	// encode/decode for sending/receiving JSON to/from a stream
	// it automatically decode JSON to GO and store data at memory adderss of the &u (VARIABLE) of type User struct, means automatically stored in struct( struct has json tags)
	json.NewDecoder(req.Body).Decode(&u) // JSON to GO

	// change Id of User struct
	u.Id ="007"

	// checking
	fmt.Println(u)

	// Again Marshal into JSON
	ujson, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
	}

	// write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", ujson)
}


func deleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: write code to delete user
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Write code to delete user\n")
}

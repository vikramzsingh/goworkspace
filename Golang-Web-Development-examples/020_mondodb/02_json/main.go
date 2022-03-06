package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"fmt"
	"encoding/json"
	"Golang-Web-Development/020_mondodb/02_json/models"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	// added route plus parameter
	r.GET("/user/:id", getUser) // reciving parameter through url
	http.ListenAndServe(":8080", r)
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	s := 	`<!DOCTYPE html>
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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK) // 200
	w.Write([]byte(s)) // for writing data on web browser // Header payload takes data in the form slice of byte []byte, the write on browser
}

func getUser(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintln(w, p.ByName("id")) // variable name in url of getUser

	// create data
	u := models.User{
		Name:   "James",
		Gender: "Male",
		Age:    54,
		Id:     p.ByName("id"),
	}

	// Marshal into JSON or / GO to JSON
	uj, err := json.Marshal(u) // uj -> user-json (JSON OBJECT/DATA)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(uj)

	// write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}
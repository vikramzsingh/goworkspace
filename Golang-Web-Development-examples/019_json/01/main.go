package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

)

type person struct {
	Fname string
	Lname string
	Items []string
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/mrshl", mrshl)
	http.HandleFunc("/encod", encod)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	s := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>FOO</title>
	</head>
	<body>
	You are at foo
	</body>
	</html>`

	io.WriteString(w, s)
}

// with Marshal function
func mrshl(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"shooting", "running", "eating"},
	}

	// Go to JSON
	json, err := json.Marshal(p) // returns []byte, error
	if err != nil {
		log.Println(err)
	}
	fmt.Println(json)

	//io.WriteString(w, string(json))
	w.Write(json) // automatically converted to string in web browser
}

// with NewEncoder function
func encod(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"gun", "guns", "only guns"},
	}

	// Automatically Go converted into JSON, then written on browser with io writer
	err := json.NewEncoder(w).Encode(p) // returns error
	if err != nil {
		log.Println(err)
	}
}

package main

import (
	"log"
	"os"
	"text/template"
)

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person        // embedded structure (person struct is embedded)
	LicenseToKill bool
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p1 := doubleZero{ // p1 is a value of type doubleZero
		person: person{ // Inner type person get promoted to outer type doubleZero
			Name: "James Bond",
			Age:  34,
		},
		LicenseToKill: true,
	}

	err := tpl.Execute(os.Stdout, p1)
	if err != nil {
		log.Fatalln(err)
	}
}

package main

import (
	"log"
	"os"
	"text/template"
)

type sage struct {
	Name  string
	Motto string
}

var tpl4 *template.Template

func init() {
	tpl4 = template.Must(template.ParseFiles("tpl4.gohtml"))
}
func main() {

	buddha := sage{
		Name:  "Buddah",
		Motto: "The belief of no belifes",
	}

	err := tpl4.ExecuteTemplate(os.Stdout, "tpl4.gohtml", buddha)
	if err != nil {
		log.Fatalln(err)
	}
}

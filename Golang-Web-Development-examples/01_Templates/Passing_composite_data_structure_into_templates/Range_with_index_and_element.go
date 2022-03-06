package main

import (
	"log"
	"os"
	"text/template"
)

var tpl1 *template.Template

func init() {
	tpl1 = template.Must(template.ParseFiles("tpl1.gohtml"))
}

func main() {
	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "shiva"}
	err := tpl1.ExecuteTemplate(os.Stdout, "tpl1.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}

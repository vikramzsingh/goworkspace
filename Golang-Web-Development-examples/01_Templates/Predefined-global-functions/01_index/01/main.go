package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl := template.Must(template.ParseFiles("tpl.gohtml"))

	xs := []string{"zero", "one", "two", "three", "four", "five", "six"}

	err := tpl.Execute(os.Stdout, xs)
	if err != nil {
		log.Fatalln()
	}
}

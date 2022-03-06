package main

import (
	"log"
	"os"
	"text/template"
)

type sage1 struct {
	Name  string
	Motto string
}

var tpl5 *template.Template

func init() {
	tpl5 = template.Must(template.ParseFiles("tpl5.gohtml"))
}
func main() {
	gandhi := sage1{
		Name:  "Gandhi",
		Motto: "Mahatma Gandhi.....",
	}

	buddha := sage1{
		Name:  "Buddha",
		Motto: "Belief of life.",
	}

	shiva := sage1{
		Name:  "Shiva",
		Motto: "Lord.....",
	}

	balmiki := sage1{
		Name:  "Balmiki",
		Motto: "Balmiki....",
	}

	// slice of sage1 struct
	// added multiple sturct into slice of sage1
	sages := []sage1{gandhi, buddha, shiva, balmiki}

	err := tpl5.ExecuteTemplate(os.Stdout, "tpl5.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}

// Variables_in_templates video 10

package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", `Relese self-focus; embrace other-focus`)
	if err != nil {
		log.Fatalln(err)
	}
}

// you can also store values in variables inside template like php, and can be able to use in template

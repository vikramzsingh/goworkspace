package main

import (
	"log"
	"os"
	"text/template"
)

var tpl3 *template.Template

func init() {
	tpl3 = template.Must(template.ParseFiles("tpl3.gohtml"))
}
func main() {
	// sages := map[string]string{"india":"Gandhi", "America":"MLK", "Meditate":"Buddha", "lord":"shiva"}
	sages := map[string]string{
		"india":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"lord":     "shiva",
	}

	err := tpl3.ExecuteTemplate(os.Stdout, "tpl3.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}

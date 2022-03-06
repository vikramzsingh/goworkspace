package main

import (
	"log"
	"os"
	"text/template"
)

var tpl2 *template.Template

func init() {
	tpl2 = template.Must(template.ParseFiles("tpl2.gohtml"))
}
func main() {
	// sages := map[string]string{"india":"Gandhi", "America":"MLK", "Meditate":"Buddha", "lord":"shiva"}
	sages := map[string]string{
		"india":    "Gandhi",
		"America":  "MLK",
		"Meditate": "Buddha",
		"lord":     "shiva",
	}

	err := tpl2.ExecuteTemplate(os.Stdout, "tpl2.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
}

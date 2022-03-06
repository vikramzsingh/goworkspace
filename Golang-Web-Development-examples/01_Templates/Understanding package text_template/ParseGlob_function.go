// ParseGlob function, parse into templates folder in the project directory(folder)

package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// Syntax: ParseGlob function:-
	// 		func ParseGlob(pattren string) (*Template, error)

	// we going to look in template folder, and parse all files with this extension .gnao
	// tpl, err := template.ParseGlob("templates/*.gnao")
	tpl, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, nil) // para takes: (io.writer, data); returns: error
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}

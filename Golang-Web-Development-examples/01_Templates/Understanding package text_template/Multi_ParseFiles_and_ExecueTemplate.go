package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("one.gnao") // returns: *Template, error
	if err != nil {
		log.Fatalln(err)
	}

	// Execute this Template
	// Return: error
	err = tpl.Execute(os.Stdout, nil) // os.Stdout prints data on Console
	if err != nil {
		log.Fatalln(err)
	}

	// Multiple files for Parse or adding multiple template files into above tpl(VAAIABLE)
	tpl, err = tpl.ParseFiles("two.gnao", "vespa.gnao") // returns: *Template, error
	if err != nil {
		log.Fatalln(err)
	}

	// Execute one template out of multiple Templates
	err = tpl.ExecuteTemplate(os.Stdout, "vespa.gnao", nil)
	if err != nil {
		log.Println(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gnao", nil)
	if err != nil {
		log.Println(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gnao", nil)
	if err != nil {
		log.Println(err)
	}

	// if we have multiple template then Execute function, Execute the first template
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Println(err)
	}
}

// Syntax: of Execute function:-
// 	func (t *Template) ExecuteTemplate(wr io.Write, name string, data interface{}) error

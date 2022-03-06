package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var tpl *template.Template //pointer to template package to Template SOME-TYPE/VARIABLE/FUNCTION
// package template Type Template
// its built-in inside text/template package

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml")) // return: *Template
}

// func init(), initilization program once. when program starting up it runs
// when program starts up we are going to pass ParseGlob as parameter because
// ParseGlob returns: (*Template, error) and Must function takes: *Template, error
// so we directly passing ParseGlob function.
// NOTE:- Must function also check error for us

func main() {
	fmt.Printf("Type of tpl VARIABLE is %T\n", tpl)
	err := tpl.Execute(os.Stdout, nil)
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

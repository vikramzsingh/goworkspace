//Understanding package text/template: parsing a file // Video 8
// parsing means we have bunch of files. We bring these files into our program and use them.
// Here assume it is a container, into which  we parse our template

package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// tpl store the data written in tpl.gohtml file. // fetch data from tpl.gohtml file
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	// Excute prints data as it has os.stdout, as os.stdout prints data on console
	err = tpl.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln(err)
	}
}

// Syntax: of ParseFiles function:-
// 		func ParseFiles(filenames ...string) (*Template, error)
// it has unlimited number strings of filesnames, because ...string (Varaidic parameter)
// e.g:
// 		tpl, err := template.ParseFiles("tpl.gohtml", "tpl.gohtml")
// it returns printer to the Template and error (*Template, error)

// we assume the returned pointer to template(tpl VARIABLE) is a container holding all the template which is parsed

// Syntax: of Execute function:-
// 		func (t *Template) Execute(wr io.Write, data interface{}) error
// it takes Write, where we going to write in file, and data
// return error

package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml") // returns: *Template, error
	if err != nil {
		log.Fatalln(err) // stop program with exit code 1
	}
	// tpl holds pointer to template file(tpl.gohtml)

	// create new file to store template file data
	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer nf.Close()

	// Execute the template by inserting the tpl.gohtml file data into index.html
	// tpl.Execute indicate which template is going to execute, as tpl holds pointer to tpl.gohtml file
	// Execute automatically write data in index.html, because of tpl.Execute
	// here we not passing any data, only doing nil at data interface{} Parameter
	err = tpl.Execute(nf, nil) //takes: new file(nf), data; and return: error
	if err != nil {
		log.Fatalln(err)
	}
}

// Syntax: of Execute function:-
// 		func (t *Template) Execute(wr io.Write, data interface{}) error
// the writer interface, write data automatically in the file

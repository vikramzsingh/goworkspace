// Parsing_data_into_templates video 9

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
	// Execute template with data passing
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}

// Syntax for displaying data in html/gohtml is curly braces to curly braces {{.}}, and data is (.) dot
// (.) dot is the current value of the data at that point in execution.
// if i have just passed in a piece of data it's that data that i have just passed in.
// if that data slice and if we ranged over slice the out of range is new input data of the next field.

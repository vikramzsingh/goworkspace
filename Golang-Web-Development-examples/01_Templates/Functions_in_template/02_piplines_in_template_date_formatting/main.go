package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

// create a FuncMap to register function for templates
// FuncMap is built-in map -->type FuncMap = map[string]interface{}
// type is template.FuncMap with composite literal{} in text/template package
var fm = template.FuncMap{
	"fdateMDY": monthDateYear,
}

func monthDateYear(t time.Time) string {
	return t.Format("01-02-2006") // func (t Time) Format(layout string) string
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}

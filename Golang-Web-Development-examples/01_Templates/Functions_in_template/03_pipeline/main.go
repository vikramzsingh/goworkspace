// Pipelining:- The output of one thing is the input of the next thing is called Pipelining
package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func double(x int) int {
	return x + x
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}

func sqroot(x float64) float64 {
	return math.Sqrt(x)
}

// create a FuncMap to register functions
// FuncMap is built-in text/template package
// type FuncMap = map[string]interface{}
var fm = template.FuncMap{
	"fdbl":  double,
	"fsq":   square,
	"fsqrt": sqroot,
}

func main() {
	f, err := os.Create("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	err = tpl.ExecuteTemplate(f, "tpl.gohtml", 3)
	if err != nil {
		log.Fatalln(err)
	}

}

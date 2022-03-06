// To pass func into our template, we need some data structure that aggregate together the functions we passing
// That data structure is map, infact there is special type of map is called FuncMap (from package text/template)
// type FuncMap = map[string]interface{} // empty-interface

package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

// Create a FuncMap to register functions. (FuncMap is built-in from text/template package)
// "uc" is what the func will be called in the template (with uc func is called in template)
// "uc" is the ToUpper func from package strings (ToUpper is built-in func)
// "ft" is func I declared
// "ft" slices a string, returning the first three characters
var fm = template.FuncMap{ // type (template.FuncMap is type) with composite literal
	"uc": strings.ToUpper, // built-in func from strings package
	"ft": firstThree,      // User-defined function registering to passing into templates
	// Note:- key is string, value is function
}

// In New func we are not giving any name/string
// Because we are not going to parse
// Then we attach our functions before we Parse our files
func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
	// Funcs() is built-n function in text/template package
	// func (t *Template) Funcs(funcMap FuncMap) *Template
	// It takes parameter of type FuncMap (funcMap is just a IDENTIFIRE/variable to store)
}

// User-defined function
func firstThree(s string) string {
	s = strings.TrimSpace(s) // removing before and after space from sting
	s = s[:3]
	return s
}

func main() {
	buddha := sage{Name: "Buddha", Motto: "Nature is life"}
	gandhi := sage{
		Name:  "Gandhi",
		Motto: "Salt Revolution",
	}

	t := car{Manufacturer: "Tata", Model: "XUV 300", Doors: 4}
	m := car{
		Manufacturer: "Mahindra",
		Model:        "Scorpio",
		Doors:        4,
	}
	h := car{
		Manufacturer: "Honda",
		Model:        "Honda city",
		Doors:        4,
	}

	sages := []sage{buddha, gandhi}
	cars := []car{t, m, h}

	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		Wisdom:    sages,
		Transport: cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}

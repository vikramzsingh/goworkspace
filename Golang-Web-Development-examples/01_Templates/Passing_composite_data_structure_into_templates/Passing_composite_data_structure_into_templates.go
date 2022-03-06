// Passing_composite_data_structure_into_templates

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
	/*
		nf, err := os.Create("index.html")
		if err != nil {
			log.Fatalln(err)
		}
	*/

	sages := []string{"Gandhi", "MLK", "Buddha", "Jesus", "shiva"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", sages)
	if err != nil {
		log.Fatalln(err)
	}
	/*
		f, err := os.Open("index.html")
		if err != nil {
			log.Fatalln(err)
		}

		bs, err := ioutil.ReadAll(f)
		if err !=  nil {
			log.Fatalln(err)
		}

		fmt.Println(bs)
		fmt.Println("The data:\n", string(bs))

	*/
}

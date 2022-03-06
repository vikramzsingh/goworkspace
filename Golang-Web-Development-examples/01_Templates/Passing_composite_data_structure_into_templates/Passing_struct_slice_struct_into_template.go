package main

import (
	"log"
	"os"
	"text/template"
)

var tpl6 *template.Template

type sage2 struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

type item struct {
	Wisdom    []sage2
	Transport []car
}

func init() {
	tpl6 = template.Must(template.ParseFiles("tpl6.gohtml"))
}
func main() {
	b := sage2{
		Name:  "Gandhi",
		Motto: "Mahatma Gandhi.....",
	}

	g := sage2{
		Name:  "Buddha",
		Motto: "Belief of life.",
	}

	m := sage2{
		Name:  "Shiva",
		Motto: "Lord.....",
	}

	t := car{
		Manufacturer: "TATA",
		Model:        "XUV",
		Doors:        4,
	}

	c := car{
		Manufacturer: "Toyota",
		Model:        "Corolla",
		Doors:        4,
	}

	sages1 := []sage2{b, g, m}
	cars := []car{t, c}

	data := item{
		Wisdom:    sages1,
		Transport: cars,
	}

	err := tpl6.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

}

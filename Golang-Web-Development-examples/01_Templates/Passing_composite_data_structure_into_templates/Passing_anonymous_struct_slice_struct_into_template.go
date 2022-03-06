package main

import (
	"log"
	"os"
	"text/template"
)

var tpl7 *template.Template

type sage3 struct {
	Name  string
	Motto string
}

type car1 struct {
	Manufacturer string
	Model        string
	Doors        int
}

func init() {
	tpl7 = template.Must(template.ParseFiles("tpl7.gohtml"))
}
func main() {
	s := sage3{
		Name:  "Shiva",
		Motto: "Om namo shivaya",
	}

	j := sage3{
		Name:  "James",
		Motto: "Bond James bond",
	}

	c := sage3{
		Name:  "Caption America",
		Motto: "Frozen in the Ice",
	}

	m := car1{
		Manufacturer: "Mahindra",
		Model:        "Xuv 500",
		Doors:        4,
	}

	t := car1{
		Manufacturer: "Tata",
		Model:        "Strome",
		Doors:        4,
	}

	sages3 := []sage3{s, j, c}
	cars1 := []car1{m, t}

	data := struct { // Anonymous struct
		Wisdom    []sage3
		Transport []car1
	}{
		Wisdom:    sages3,
		Transport: cars1,
	}

	err := tpl7.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}

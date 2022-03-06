package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

/*
	or
type course struct {
	Number string
	Name string
	Units string
}
*/

type semester struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	y := year{
		Fall: semester{
			Term: "Fall",
			Courses: []course{ // slice of course struct
				course{Number: "CSCI-40", Name: "Programming in Go", Units: "4"},
				course{"CSCI-130", "Web Programming in Go", "4"}, // course struct
				course{"CSCI-140", "Mobile App with Go", "4"},    // course struct
			},
		},

		Spring: semester{
			Term: "Spring",
			Courses: []course{ // slice of course struct
				course{ // couser struct
					Number: "CSCI-50",
					Name:   "Advanced Go",
					Units:  "5",
				},
				course{Number: "CSCI-190", Name: "Advanced Web", Units: "5"}, // course struct
				course{"CSCI-191", "Advanced Mobile Go", "5"},
			},
		},
	}

	err := tpl.Execute(os.Stdout, y)
	if err != nil {
		log.Fatalln(err)
	}
}

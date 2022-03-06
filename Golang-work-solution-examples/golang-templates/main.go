package main

import (
	"fmt"
	"text/template"
)

func main() {
	val, err := template.New("html-tmpl").Parse("Hello world")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)	
}

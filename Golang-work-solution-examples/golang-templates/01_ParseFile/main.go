package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

// Prints out the template without passing any value using the text/template package
func main() {
    template, err := template.ParseFiles("template-01.txt")
    // Capture any error
    if err != nil {
        log.Fatalln(err)
    }
    // Print out the template to std
    template.Execute(os.Stdout, nil)
	fmt.Println()
	template.Execute(os.Stdout, 6)
}
//OUTPUT
// Hi <no value>
// You are welcome to this tutorial
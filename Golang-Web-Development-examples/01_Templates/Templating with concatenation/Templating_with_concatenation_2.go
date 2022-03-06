package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// this code not working
func main() {
	name := os.Args[1] // creating and getting two variables from os.args
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

	// string of html document
	str := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World</title>
	</head>
	<body>
	<h1>` + name + `</h1> 
	</body>
	</html>
	`)

	nf, err := os.Create("index-file-1.html")
	if err != nil { // means error reported
		log.Fatalln(err)
	}
	defer nf.Close()

	// putting tpl(Templte data) in index-file.html
	// nf is the file which need to be write in
	// tpl is data which is going to index-file.html

	//s := strings.NewReader(tlp) // NewReader(built-in function) reading byte data from tpl, and strings(built-in function) converting byte data into string
	io.Copy(nf, strings.NewReader(str))
}

package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Todd Mcleod" // string Todd Mcleod

	// string of html document
	tlp := `
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
	`
	fmt.Println(tlp)

	nf, err := os.Create("index-file.html")
	if err != nil { // means error reported
		log.Fatalln(err)
	}
	defer nf.Close()

	// putting tpl(Templte data) in index-file.html
	// nf is the file which need to be write in
	// tpl is data which is going to index-file.html

	//s := strings.NewReader(tlp) // NewReader(built-in function) reading byte data from tpl, and strings(built-in function) converting byte data into string
	io.Copy(nf, strings.NewReader(tlp))
}

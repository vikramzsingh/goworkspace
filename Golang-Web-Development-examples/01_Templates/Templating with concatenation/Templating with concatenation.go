// Templating_with_concatenation

package main

import "fmt"

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
}

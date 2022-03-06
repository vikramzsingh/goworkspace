package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Application is started...")

	// flag.Bool("name of flag/ by which flag is activated", "default value of flag", "purpose OR defination of flag")
	output := flag.Bool("output", false, "If ture allowed to perform action, if false not allowed to perform action")

	// let we have a flag, that represents a input file,
	// let default value for this flag would be file.csv,
	// last parameter will tell what this flag do OR work of this flag
	input_file_flag := flag.String("input", "file.csv", "The path to the input file")

	flag.Parse() // In order to actuallay use flag we need to parse our flag

	
	if (*output) { // value is bool type
		fmt.Println(output)
		fmt.Println(&output)
		fmt.Println(*output)
	}

	if *input_file_flag != "file.csv" {
		fmt.Println(*input_file_flag)
	}
}

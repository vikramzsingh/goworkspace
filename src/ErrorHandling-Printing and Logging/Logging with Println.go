package main

// log.Println() allow us to send or write error details into files.

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	log.SetOutput(f) // takes io writer and return log

	f2, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened", err)
		log.Println("error happened", err) // This log error information going to log.SetOutput(f), f is log.txt file, so log error Information will save into log.txt file
		//log.Fatalln(err)
		//panic(err)
	}
	defer f2.Close()

	fmt.Println("check log.txt file in the directory")

}

// Println calls Output to print to the  standard logger. Arguments are handled in the manner of fmt.Println

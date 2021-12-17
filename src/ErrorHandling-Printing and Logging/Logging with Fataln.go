package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	defer foo() // with log.Fataln. The program terminates immediately, deferred function are not run.
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened", err)
		//log.Println("error happened", err)
		log.Fatalln(err)
		//panic(err)
	}
}

func foo() {
	fmt.Println("When os.Exit() is called, deffered function don't run")
}

//Fataln function calls os.Exit(1), means close the program.
//Fatal is equivalent to Println followed by a call os.Exit(1)
// Exit causes the current program to exit with the given status code. Conventionally, code zero indicates success, non-zero an error. The program terminates immediately, deferred function are not run.

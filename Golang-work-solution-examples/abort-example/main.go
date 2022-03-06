package main

import (
	"fmt"
)

func main() {	
	if foo() == "foo" {
		fmt.Println("ready to return, not executing further")
		return
	}
	boo("boo")
	bar()
}

func foo() string {
	return "foo"
	// return "not foo"
}

func boo(str string) {
	if str == "boo" {
		fmt.Println("buu not executing further")
		return
	}
	fmt.Println("End of buu")
}

func bar() {
	fmt.Println("bar")
}
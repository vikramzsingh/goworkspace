// The-var-keyword video 28
package main

import (
	"fmt"
)

var y = 43 // globle variable are declared by var-keyword
// scope is globle

var z int //syntax => var-Keyword variable-name int-type(var z int)
// z stores 0 value as int default value is 0(Zero), bool default value is false
func main() {
	x := 42 // local variable(only declarable within the code body)
	// scope is local
	fmt.Println(x)

	//var y = 43 // Another way of declaring variable with var-Keyword
	fmt.Println(y)

	foo() // function calling

	fmt.Println(z)

	var s int // method for creating Enpty-variable
	s = 54
	fmt.Println(s)
}

// function declaration
func foo() {
	fmt.Println("again:", y)
}

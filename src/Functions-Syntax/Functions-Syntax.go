// Functions_Syntax video 107
// syntax:-
// function(receiver) identifir(parameters) (returns) { code... }
package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
	foo()                  // calling foo function // here we pass arguments if has any
	bar("vikram")          //passing arg
	s1 := woo("moneypeny") // calling woo
	fmt.Println(s1)
	x, y, z := mouse("Ian", "Fleming") // multiple returns
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

// func (r  reciver) identifier(parameter(s)) (return(s)) { code... }
// Everything in Go is PASS BY VALUE

// takes no argument, has no parameters and returns notthing
func foo() { // para defines here
	fmt.Println("Hello, Vikram function foo")
}

// with para defined, returns notthing
func bar(s string) {
	// The scope of this s In para is only whith in this function
	fmt.Println("Hello,", s, "function bar")
}

func woo(st string) string {
	//s := fmt.Sprint("Hello form woo, ", st)
	//fmt.Println("By s:",s)
	//return s

	return fmt.Sprint("Hello form woo, ", st) // Sprint() stores the given data in string format, we store this data in variable then we canm print it.
}

// multiple return function
func mouse(fn string, ln string) (string, bool, bool) { // Three returns
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := true
	c := false
	return a, b, c // returning string, bool
}

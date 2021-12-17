// Functions_Returning_a_func

package main

import "fmt"

func main() {

	//x := bar() // here bar() is returning func() int, and we storing it in x,

	//fmt.Printf("%T\n", x)

	// as x storing anonymous function expression, hence we can run it.

	//i := x() // x stores anonymous func expression
	//fmt.Println(i)
	//or
	//fmt.Println(x())
	fmt.Println(bar()())

	// bar() returns func() int [Anonymous func expression], then we execute it with ().
}

// function bar() is returning ( func() int ).
// And anonymous function, func() is returning int
func bar() func() int {
	return func() int {
		return 451
	}
}

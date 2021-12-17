// Functions_Defer video 110
// defer is a KEYWORD
// difer KEYWORD will differ the execution of function

package main

import "fmt"

func main() {
	// defer KEYWORD evaluate foo() but it will execute after the end of this code block ends
	defer foo() //with difer KEYWORD, Right where we opening the function/file we also closeing that function/file
	// but the output wil be printer at the end of function body block, where it is called
	bar()
	boo()
} // after exiting body, at this point the function difr foo() display output

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func boo() {
	fmt.Println("boo")
}

//A "defer" statement invokes a function whose execution is deferred to the moment the surrounding function returns, either because the surrounding function executed a return statement, reached the end of its function body, or because the corresponding goroutine is panicking.

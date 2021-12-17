//semi-colon

package main

import "fmt"

func main() {
	// In Go-lang we use semi-colon to a statement
	// Short-declaration is statement, so we put semi conlon
	if x := 42; x == 42 { // you can also declare x here at condition.
		fmt.Println("001")
	}

	// In Go-lang, In a statement complier puts semi-colon automatically at the end
	// Or if we want 2(Two) statements in one line. So we put semi-colon after first statement, example given below:
	fmt.Println("here's a statement")
	fmt.Println("something else")

	fmt.Println("\n")
	fmt.Println("here's a statement") // complier put semi-colon automatically
	fmt.Println("something else")
}

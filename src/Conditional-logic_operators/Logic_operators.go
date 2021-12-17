// Conditional_logic_operators video

package main

import "fmt"

func main() {
	fmt.Printf("true && true\t %T\n", true && true) //%T printing type
	fmt.Printf("true && true\t %v\n", true && true) //%v printing value(which is true)
	fmt.Printf("true && false\t %v\n", true && false)
	fmt.Printf("true || true\t %v\n", true || true)
	fmt.Printf("true || false\t %v\n", true || false)
	fmt.Printf("!true \t\t\t %v\n", !true)
	fmt.Printf("!false \t\t\t %v\n", !false)
}

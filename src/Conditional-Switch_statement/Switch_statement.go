//Conditional_Switch_statement video 64

package main

import "fmt"

func main() {
	// Missing switch expression is default true
	switch {
	case false:
		fmt.Println("This should not print")
	case (2 == 4):
		fmt.Println("This should not print2")
	case (3 == 3): // Note: Only first true case prints
		fmt.Println("prints")
		fallthrough // with this you can run next case(independent of true or false)
	case (4 == 4):
		fmt.Println("also true, does it print?")
		fallthrough
	case (7 == 9):
		fmt.Println("Not true1")
		fallthrough
	case (11 == 14):
		fmt.Println("Not true2")
		fallthrough
	case (15 == 15):
		fmt.Println("true15")
		fallthrough
	default:
		fmt.Println("This is default")
	}
}

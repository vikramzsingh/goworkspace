// Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.

package main

import "fmt"

func main() {
	//var favSport string = "Football"
	favSport := "Football"

	switch favSport {
	case "Hockey":
		fmt.Println("This Should print Hockey.")
	case "Football":
		fmt.Println("This Should print Football.")
	case "Skiing":
		fmt.Println("This Should print Skiing.")
	case "Surfing":
		fmt.Println("This Should print Surfing.")
	default:
		fmt.Println("This is default.")
	}
}

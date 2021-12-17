package main

import "fmt"

func main() {
	switch "Bond" {
	case "Monneypenny", "Bond", "Doom":
		fmt.Println("miss monney or bond or doom")
	case "M":
		fmt.Println("m")
	case "Q":
		fmt.Println("Q")
	default:
		fmt.Println("this is default")
	}
}

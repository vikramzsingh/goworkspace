package main

import "fmt"

// Calculate return x + 2
func Calculate(x int) int {
	result := x + 2
	return  result
}

// Add return x + y
func Add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println("Hello testing")
	result := Calculate(2)
	fmt.Println(result)
}

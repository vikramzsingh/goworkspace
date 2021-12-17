// Functions_Closure

package main

import "fmt"

func main() {
	var x int
	fmt.Println(x)

	// Clouser liniting the access of code
	{
		y := 42
		fmt.Println(y)
	}
	//fmt.Println(y)

	x++
	fmt.Println(x)
}

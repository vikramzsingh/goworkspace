package main

import "fmt"

func main() {
	fmt.Println(Incrementor()())

	// or

	x := Incrementor()
	fmt.Printf("%T\n", x)
	i := x()
	fmt.Println(i)
	fmt.Println()
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}

func Incrementor() func() int {
	var y int
	return func() int {
		y++
		return y
	}
}

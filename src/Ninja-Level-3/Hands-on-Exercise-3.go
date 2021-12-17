/*
Create a for loop using this syntax
for condition { }
Have it print out the years you have been alive.
*/

package main

import "fmt"

func main() {
	// Multi Variable declaration :-
	// x := 1998
	// count := 0
	x, count := 1998, 0
	for x <= 2021 {
		fmt.Println(x)
		x++
		count++
	}
	fmt.Println("Number of years:", count)
}

/*
Create a for loop using this syntax
for { }
Have it print out the years you have been alive.
*/
package main

import "fmt"

func main() {
	x := 1998
	for {
		if x <= 2021 {
			fmt.Println(x)
		}
		x++
	}
}

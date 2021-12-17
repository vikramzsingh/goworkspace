// Slice_Multi_dimensional_slice video 84

package main

import "fmt"

func main() {
	//A composite literal is created by having the TYPE followed by CURLY BRACES and then putting the appropriate values in the curly brace area.
	// create a slice with composite literal
	// i.e TYPE{} --> []string is TYPE, and followed by {}
	jb := []string{"James", "Bond", "chocolate", "martini"}
	fmt.Println(jb)

	mp := []string{"Miss", "Moneypeny", "strawberry", "Hazenut"}
	fmt.Println(mp)

	// slice of slice of string, this is multi-dimentional slice
	xp := [][]string{jb, mp} //slice of slice of string
	//xp := [][]string{{"aaaa", "bbbb"},{"DDDD", "EEEE"}}
	fmt.Println(xp)
}

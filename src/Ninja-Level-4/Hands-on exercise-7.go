// Hands-on exercise-7 video 94
/*
Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "Helloooooo, James."

Range over the records, then range over the data in each record.

*/
package main

import "fmt"

func main() {
	s1 := []string{"James", "Bond", "Shaken, not stirred"}
	s2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(s1)
	fmt.Println(s2)

	s := [][]string{s1, s2} // slice of a slice of string ([][]string)
	fmt.Println(s)

	for i, v := range s {
		fmt.Println("record: ", i, v)
		for j, v1 := range v {
			fmt.Printf("\t Index: %v \t value: \t %v \n", j, v1)
		}
	}
	fmt.Println("Print With only values:-")
	// with only values
	for _, v := range s { // _ it throws away the first value, i.e index
		fmt.Println(v)
		for _, v1 := range v {
			fmt.Println("\t value: \t", v1)
		}
	}
}

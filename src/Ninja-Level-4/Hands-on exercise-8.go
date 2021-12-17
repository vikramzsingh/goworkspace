// Hands-on exercise-8 video 95
/*
Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things. Store three records in your map. Print out all of the values, along with their index position in the slice.

	`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
`no_dr`, `Being evil`, `Ice cream`, `Sunsets`

*/

package main

import "fmt"

func main() {
	// map[Key_Type]Value_Type{
	// 		....
	// }
	// here Key_type is:- string
	//		Value_Type is:- []string

	// The map type is:- map[string][]string
	m := map[string][]string{
		// because Value_Type is []string (slice of string)
		"bond_james": []string{`Shaken, not stirred`, `Martinis`, `Women`},

		`moneypenny`: []string{`James Bond`, `Literature`, `Computer Science`},

		`no_dr`: []string{"Being evil", "Ice cream", "Sunsets"},
	}
	// to check for given key is there any value
	if v, ok := m["bond_james"]; ok {
		fmt.Println("Status of comma ok idiom:", ok)
		fmt.Println("Value:", v)
		for i, val := range v {
			fmt.Printf("\t Index %v \t value \t %v \n", i, val)
		}
	}
	fmt.Println("\n\n")
	// To Print out all of the values, along with their index position in the slice.
	for k, val := range m {
		fmt.Println("This is record for key:", k)
		for i, val1 := range val {
			fmt.Printf("\t Index: %v \t value: \t %v \n", i, val1)
		}
	}
}

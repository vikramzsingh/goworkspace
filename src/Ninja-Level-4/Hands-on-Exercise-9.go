// Hands on exercise-9 video 96
/*
Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop
*/
package main

import "fmt"

func main() {
	// create a map
	m := map[string][]string{
		// because Value_Type []string (slice of string)
		"bond_james": []string{`Shaken, not stirred`, `Martinis`, `Women`},

		`moneypenny`: []string{`James Bond`, `Literature`, `Computer Science`},

		`no_dr`: []string{"Being evil", "Ice cream", "Sunsets"},
	}

	fmt.Println(m)

	// adding key and value to map

	m["vikram"] = []string{"B.Tech", "Learning Go", "CSE"}
	fmt.Println(m)

	//print the map out using the “range” loop
	for k, v := range m {
		fmt.Printf("key: %v \t\t value: %v\n", k, v)
		for i, val := range v {
			fmt.Printf("\tIndex: %v \t value: %v\n", i, val)
		}
	}
}

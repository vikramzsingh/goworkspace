// JSON_Marshal

package main

import (
	"encoding/json"
	"fmt"
)

type person struct { // In JSON Marshal fields first letter must be capital
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Vikram",
		Last:  "singh",
		Age:   23,
	}

	p2 := person{
		First: "James",
		Last:  "bond",
		Age:   45,
	}

	people := []person{p1, p2} // slice of person
	//var people []person
	//people = []person{p1, p2}
	fmt.Println(people) // Go data

	// Syntax of Marshal JSON:
	// Marshal(v interface{}) ([]byte, error)
	// It returns a slice of byte ([]byte) and error
	// converting data into JSON data
	bs, err := json.Marshal(people) // JSON Marshal takes the value of any type, because it has empty INTERFACE
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
	// NOTE:- In JSON Marshal fields first letter must be capital
}

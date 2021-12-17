// JSON_unmarshal

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"_"` // throw away or not use it
	Age   int    `json:"Age"`
}

func main() {
	// JSON data
	s := `[{"First":"Vikram","Last":"singh","Age":23},{"First":"James","Last":"bond","Age":45}]`

	// s :=`{"First":"Vikram","Last":"singh","Age":23}`

	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	//people := []person{}
	var people []person // as structure(mixed data) type data is comming

	// var people person
	// type person has fields: First, Last, Age

	// JSON Unmarshal Syntax:
	// func Unmarshal(data []byte, v Interface{}) error
	// converting go data
	err := json.Unmarshal(bs, &people) // it must need pointer to the variable (like people) to store data.
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\nAll of the data:", people)

	for i, v := range people {
		fmt.Println("\nPerson Number:", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}

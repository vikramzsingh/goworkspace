// JSON_Encode
// encode from go to JSON
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	/*
		var p1 person
		p1 = person{
			First:       "James",
			Last:        "Bond",
			Age:         32,
			notExported: 007, //This field not going to encode into JSON, because this field is not start with capital letter
		}
	*/
	// or

	p1 := person{"James", "Bond", 32, 007}
	fmt.Println(p1)

	json.NewEncoder(os.Stdout).Encode(p1)
	// at this we get pointer to Encoder so we attach Encode method to it

}

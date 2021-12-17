// JSON_Decode

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	var p1 person // to store Go data

	// JSON OBJECT passed into NewReader
	// we converted string into Reader, as NewDecoder takes Reader
	rdr := strings.NewReader(`{"First": "James", "Last": "Bond", "Age": 20}`)
	json.NewDecoder(rdr).Decode(&p1) // JSON to go
	// NewDecoder take reader

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Println(p1.notExported)
}

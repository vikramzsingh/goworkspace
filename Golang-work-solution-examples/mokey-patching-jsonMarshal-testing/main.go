package main

import (
	"fmt"
	"mokey-patching-json-testing/helper"
)

type Person struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	p1 := Person {
		FirstName: "Vikram",
		LastName: "Singh",
	}
	s := MarshalData(p1)
	fmt.Println(s)
}

func MarshalData(v interface{}) string {
	sb, err := helper.ExecuteMarshal(v)
	fmt.Println(sb)
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(sb)
}
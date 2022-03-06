// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	s1 := []Person{
		{
			Name: "Vikram",
			Age:  24,
		},
		{
			Name: "Bobby",
			Age:  25,
		},
	}
	s2 := []Person{
		{
			Name: "Vikram",
			Age:  24,
		},
		{
			Name: "Bobby",
			Age:  25,
		},
	}
	fmt.Println(reflect.DeepEqual(s1, s2))
}

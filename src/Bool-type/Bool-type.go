// Bool Type video 40
package main

import (
	"fmt"
	"reflect"
)

var x bool // boolean type

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	// or by using Println
	fmt.Println(reflect.TypeOf(x))

	a := 7
	b := 42
	fmt.Println(a != b)
}

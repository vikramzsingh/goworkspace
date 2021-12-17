// Numeric_type video 42
package main

import "fmt"

var x int // uint(0 to 255) is un-signed where int is signed(-128 to 127)
var y float64

func main() {
	/* another method
	x := 42
	y := 42.34534
	*/
	x = 42 // x = 2.34534 not work because x is integer type and float value cannot be assigned
	y = 42.34534

	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
}

// byte for uint8
// rune for int32

// video 38

package main

import "fmt"

type vikram int // own-type or custom-type

var x3 vikram // x3 variable of vikram-type
var y3 int    // created a variable y3 of underlying-type of coustom-type of x3

func main() {
	fmt.Println(x3)
	fmt.Printf("%T\n", x3)
	x3 = 42
	fmt.Println(x3)

	y3 = int(x3) // converted vikram-type to int-type
	fmt.Println(y3)
	fmt.Printf("%T", y3)
}

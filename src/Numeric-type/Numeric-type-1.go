// signed int8 8byte(256 bits) range -128 to 127
package main

import "fmt"

var x1 int8 = -128 // int8 range 8byte(-128 to 127) so -129 not be able to store in int8
func main() {
	fmt.Println(x1)
	fmt.Printf("%T\n", x1)
}

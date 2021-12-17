// Bit_shifting video 47
// Bit-shift operators >> and <<
package main

import "fmt"

func main() {
	x := 4                         // short-operator used, assigned value is 4(IN DECIMAL)
	fmt.Printf("%d\t\t%b\n", x, x) // %d for DECIMAL %b for BINARY

	y := x << 1 // as x is 4(IN BINARY 100) shifting one bit to left that means 1000(IN DECIMAL 8) and assigned to y
	// y = y >> 1
	// y = x >> 1
	fmt.Printf("%d\t\t%b\n", y, y) // prints the value of y = 8 in DECIMAL and y = 1000 in BINARY
}

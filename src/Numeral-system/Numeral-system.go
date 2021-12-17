// Numeral_system video 44

package main

import "fmt"

func main() {
	s := "H"
	fmt.Println(s)

	// convert into slice of byte
	bs := []byte(s) // slice of byte contain ASCII value
	fmt.Println(bs) //Its type is uint8 or byte

	// Then we can access that value using index position zero of bs
	n := bs[0]
	fmt.Println(n)

	// we can check its type
	fmt.Printf("%T\n", n) // type is uint8 or byte

	// now convert it into binary and hexa-decimal
	fmt.Printf("%b\n", n) // or printing byte(uint8) into binary type
	fmt.Printf("%#X\n", n)
}

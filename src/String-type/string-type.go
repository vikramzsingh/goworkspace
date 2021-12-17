// String_type video 43

package main

import "fmt"

func main() {
	s := "Hello, playground" //(``)Back-tiks
	fmt.Println(s)

	bs := []byte(s) //([]byte) slice of byte and s is converted to byte
	// s is converted into byte
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	//byte is alice of uint8
	// it tell us []uint8 = slice of byte

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i]) // %#U for UTF-8 value or rune = int32
	}

	fmt.Println("")

	for i, v := range s { // ASCII code value prints v variable
		fmt.Println(i, v)
	}

	for i, v := range s {
		fmt.Printf("At index position %d we have hex %#x\n", i, v)
	}
}

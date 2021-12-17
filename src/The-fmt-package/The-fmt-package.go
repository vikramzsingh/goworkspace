// The-fmt-package video 31
package main

import "fmt"

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)  // I want  as integer-type
	fmt.Printf("%b\n", y)  // I want  as binary
	fmt.Printf("%x\n", y)  // I want as Hexa-decimal
	fmt.Printf("%#x\n", y) //I want as Hexa-decimal with 0x infornt of it
	y = 911
	fmt.Printf("%#x\n", y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)

	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y) // its a STRING-printf(Sprintf) used to print string  and assing that string-value to a variable
	// Sprintf convert the data in String, and then we can assign that string value to a variable
	// then we can print it
	fmt.Println(s)
	fmt.Printf("%v", y) // it prints the value of variable
}

// Constant video 45

package main

import "fmt"

// const is a keyword
const a int = 42 // we can also specify type
const b = 42.78
const c = "James Bond"

/* Another way of using const
const (
	a int = 42
	b = 42.78
	c = "James Bond"
)
*/

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}

//Create TYPED and UNTYPED constants. Print the values of the constants.

package main

import "fmt"

const (
	a     = 42 // UNTYPED constant
	b int = 43 // TYPED constant
)

const c = 44     // UNTYPED
const d int = 45 // TYPED

func main() {
	fmt.Println(a, b, c, d)
}

// Using iota, create 4 constants for the NEXT 4 years. Print the constant values.

package main

import "fmt"

const ( // iota can only be useable with const-keyword, it start with 0 and increment each time by one.
	// with new conat declaration it again start with 0 and increment by one ecah time
	a1 = 2017 + iota // 2017 + 0
	b2 = 2017 + iota // 2017 + 1
	c3 = 2017 + iota // 2017 + 2
	d4 = 2017 + iota // 2017 + 3
)

func main() {
	fmt.Println(a1)
	fmt.Println(b2)
	fmt.Println(c3)
	fmt.Println(d4)
}

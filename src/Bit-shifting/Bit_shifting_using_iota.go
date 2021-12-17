// Bit_shifting using iota
package main

import "fmt"

const (
	_   = iota             // this means throw iota not use it,& it iota  Value is 0
	kb1 = 1 << (iota * 10) // Value is 1(its iota) * 10 = 10 zeros
	mb1 = 1 << (iota * 10) // Value is 2(its iota) * 10 = 20 zeros
	gb1 = 1 << (iota * 10) // value is 3(its iota) * 10 = 30 zeros
)

func main() {
	fmt.Printf("%d\t\t\t%b\n", kb1, kb1)
	fmt.Printf("%d\t\t\t%b\n", mb1, mb1)
	fmt.Printf("%d\t\t%b\n", gb1, gb1)
}

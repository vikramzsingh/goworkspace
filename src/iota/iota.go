// iota vodeo 46

package main

import "fmt"

const (
	a = iota // also possible iota + 1 it gives (0+1)=1
	b = iota
	c = iota
	g = iota
	h = iota
	i = iota
)

/* Alos works
const(
	a = iota
	b
	c
)
*/
const ( // const-keyword reset it, now it start again 0,1,2
	d = iota
	e //= iota
	f //= iota
)

//const name  = iota
// you can't be able to use iota without const-keyword

const a1 = iota
const a2 = iota

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)

	fmt.Println(a1) // = 0
	fmt.Println(a2) // = 0

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}

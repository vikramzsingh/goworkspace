package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("hello World")
	
	a := 5
	fmt.Printf("value %d, Type %T \n", a, a)

	res := strconv.Itoa(a)
	fmt.Printf("value %s, Type %T \n", res, res)

	res1, _ := strconv.Atoi(res)
	fmt.Printf("value %d, Type %T \n", res1, res1)


}
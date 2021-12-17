package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt3(-10.23)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt3(f float64) (float64, error) {
	if f < 0 {
		ErrNorgateMath1 := fmt.Errorf("norgate math again: square root of negative number: %v", f)
		return 0, ErrNorgateMath1
	}
	return 42, nil
}

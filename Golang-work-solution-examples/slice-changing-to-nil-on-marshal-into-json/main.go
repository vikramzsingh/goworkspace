// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/json"
	"exp/utillib"
	"fmt"
)



func main() {
	// var hitdata = make([]utillib.Exp, 0)
	var hitdata []utillib.Exp
	
	fmt.Println(hitdata)
	fmt.Printf("%T\n", hitdata)
	
	sb, err := json.Marshal(hitdata)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("sb should be slice not nil:", string(sb))

	// var xs []string 
	xs := []string{} // with literal it is also got initilized.
	sb1, err := json.Marshal(xs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("sb1", string(sb1))
}

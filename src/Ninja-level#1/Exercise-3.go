//Exercise-3 video 36
package main

import "fmt"

var x1 int = 42
var y1 string = "James Bond"
var z1 bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x1, y1, z1) //In Sprintf(f including function) formatting string(%v it display value in variable and %T display-type of variable) is required
	fmt.Println(s)

}

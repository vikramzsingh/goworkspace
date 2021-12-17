// With runtime package we can check operating system and architecture of the machine running behind the GO-language
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)   // we can see the Operating system
	fmt.Println(runtime.GOARCH) // we can see the Architecture
	// running behind GO language
}

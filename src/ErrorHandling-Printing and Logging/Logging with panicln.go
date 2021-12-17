package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened", err)
		//log.Println("error happened", err)
		//log.Fatalln(err)
		log.Panicln(err)
		//panic(err)
	}
}

//panicln is equivalent to Println followed by a call panic

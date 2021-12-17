package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("name.txt")
	if err != nil {
		fmt.Println(err)
		return // if there is error leave func main and program will end, and defer close file
	}
	defer f.Close()

	r := strings.NewReader("James Bond")
	io.Copy(f, r)

	// OR
	/*
		bs := []byte("james bond testing")
		d, err := f.Write(bs)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Number of string written in file: ", d)
	*/
}

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	bs, err := ioutil.ReadAll(conn) // read in from connection, written by 01_write server
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(bs))
}

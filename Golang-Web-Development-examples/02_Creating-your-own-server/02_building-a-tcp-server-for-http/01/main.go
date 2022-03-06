package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		defer conn.Close()

		go handel(conn)
	}
}

func handel(conn net.Conn) {

	// read request
	request(conn)

	// write response
	respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln) // server readed first line data from connection and prints it out
		// Then it goes below to code to prints METHOD and URL from this line.
		if i == 0 { // 0 indicates, Means first line is readed which is at 0 index
			//func strings.Field(s string) []string
			m := strings.Fields(ln)[0] // --> give us METHOD
			// m := strings.Fields(ln)[1] // --> give us URL
			// while reading from connection first line comes with index 0.
			// And this line is a slice of string. like:- GET / http/1.1
			// At 0 index in this line it contains Method at 0 index
			// so we storing method in to m VARIABLE from 0 index of the first comming line
			fmt.Println("***METHOD", m)
		}
		if ln == "" {
			//header are done
			break
		}
		i++
	}
}

func respond(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

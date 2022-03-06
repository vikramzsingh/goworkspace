// read from connection using bufio scanner
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		defer conn.Close()

		go handle(conn) //we used go routine so our tcp server can handel more than one connection at a time
	}

}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() { // gives each token which is by default is line
		ln := scanner.Text() // server read from connection.
		fmt.Println(ln)      // After reading from connection, write back into server window
	}
	defer conn.Close()

	// we never get here
	// we have an open stream connection
	// how does the above reader knows when it's done?
	fmt.Println("Code got here.")

	// because the scanner.Scan() is in loop and return false at last line but you have wide open connection asking are you still there.
}

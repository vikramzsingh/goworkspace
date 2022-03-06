package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	// Create listener using net package
	// takes network string, address string
	// returns listener, error
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		// Accept implements the Accept method in the Listener interface;
		// it waits for the next call and returns a generic Conn
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		go handel(conn)
	}
}

func handel(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()                           // server read from connection
		fmt.Println(ln)                                // After reading from connection, shows in server window
		fmt.Fprintf(conn, "I heard you say :%s\n", ln) // writing into server, then its shows in connection window
		/*
			fmt.Println("Type your message to connection:\n")
			var s string
			fmt.Scanf("%s",&s,"\n")
			fmt.Fprintf(conn, s)
		*/
	}

	defer conn.Close()

	// we never got here
	// we have an open stream connection
	// how does the above reader konw the when its done?
	fmt.Println("code got here")
}

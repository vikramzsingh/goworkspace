// we created a tcp server, and read- write through that connection
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	// create listener
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		go handel(conn)
	}
}

func handel(conn net.Conn) {
	// Applying Deadline for connection
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("CONN TIMEOUT")
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(conn, "I heard you say: %s\n", line)
	}
	defer conn.Close()

	fmt.Println("********Code got here*******")
}

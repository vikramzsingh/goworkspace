package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	//   func Listen(network string, address string) net (Listener, error)
	// Listen ask for two strings:
	// network string, What kind of network we going to Listen ("tcp")
	// address string, What port do you want to Listen on ":8080"
	// Listen returns us (Listener, error) a Listener and error
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		// func (l *TCPListener) Accept() (Conn, error)
		//if Somebody calls us we accept,i.e by li.Accept() if somebody calls us we accept
		// Accept returns connection and  error, (conn, err)
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		io.WriteString(conn, "\nHello from TCP server\n")
		//fmt.Fprintln(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintln(conn, "%v", "well, I hope!")

		conn.Close()
	}
}

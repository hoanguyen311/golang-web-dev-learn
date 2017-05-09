package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
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

		io.WriteString(conn, "\nHello\n")
		fmt.Fprintln(conn, "How's your day?")
		fmt.Fprintf(conn, "%v", "Well, I hope\n")

		conn.Close()
	}
}

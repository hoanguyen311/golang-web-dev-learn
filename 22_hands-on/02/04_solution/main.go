package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := li.Accept()
		defer conn.Close()
		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(conn)

		for scanner.Scan() {
			text := scanner.Text()
			fmt.Println(text)
		}

		io.WriteString(conn, "I see you connected.\n")
	}
}

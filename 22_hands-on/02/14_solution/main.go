package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := li.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	i := 0
	for scanner.Scan() {
		text := scanner.Text()

		if i == 0 {
			paths := strings.Fields(text)
			fmt.Printf("Method: %s\n", paths[0])
			fmt.Printf("Path: %s\n", paths[1])
		} else {
			fmt.Println(text)
		}
		if text == "" {
			fmt.Println("END OF CONNECTION")
			break
		}
		i++
	}
	body := "Here I WRITE to the connection\n\n"
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	io.WriteString(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

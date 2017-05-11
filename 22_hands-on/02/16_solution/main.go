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
	body := `<!DOCTYPE html> <html lang="en"> <head> <meta charset="UTF-8"> <meta name="viewport" content="width=device-width, initial-scale=1.0"> <meta http-equiv="X-UA-Compatible" content="ie=edge"> <title>HOLY COW THIS IS LOW LEVEL</title> </head> <body> <h1>HOLY COW THIS IS LOW LEVEL</h1> </body> </html>`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	io.WriteString(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

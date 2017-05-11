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
			xs := strings.Fields(text)
			method := xs[0]
			path := xs[1]
			if path == "/" && method == "GET" {
				index(conn)
			} else if path == "/apply" && method == "GET" {
				apply(conn)
			} else if path == "/apply" && method == "POST" {
				applyProcess(conn)
			}
		}
		if text == "" {
			fmt.Println("END OF CONNECTION")
			break
		}
		i++
	}

}

func apply(conn net.Conn) {
	body := `<!DOCTYPE html> <html lang="en"> <head> <meta charset="UTF-8"> <meta name="viewport" content="width=device-width, initial-scale=1.0"> <meta http-equiv="X-UA-Compatible" content="ie=edge"> <title>Apply</title> </head> <body> <h1>APPLY</h1><form action="/apply" method="post"><button type="submit">Apply process<button></form> </body> </html>`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	io.WriteString(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func index(conn net.Conn) {
	body := `<!DOCTYPE html> <html lang="en"> <head> <meta charset="UTF-8"> <meta name="viewport" content="width=device-width, initial-scale=1.0"> <meta http-equiv="X-UA-Compatible" content="ie=edge"> <title>index</title> </head> <body> <h1>INDEX</h1> </body> </html>`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	io.WriteString(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func applyProcess(conn net.Conn) {
	body := `<!DOCTYPE html> <html lang="en"> <head> <meta charset="UTF-8"> <meta name="viewport" content="width=device-width, initial-scale=1.0"> <meta http-equiv="X-UA-Compatible" content="ie=edge"> <title>Apply</title> </head> <body> <h1>APPLY PROCESS</h1> </body> </html>`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	io.WriteString(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

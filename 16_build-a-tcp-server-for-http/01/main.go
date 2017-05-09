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
		log.Fatal(err)
	}

	for {
		conn, err := li.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	request(conn)
	response(conn)
}

func request(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	i := 0
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			m := strings.Fields(ln)[0]
			path := strings.Fields(ln)[1]
			fmt.Println(path)
			fmt.Println("***METHOD", m)
		}
		if ln == "" {
			break
		}
		i++
	}
}
func response(conn net.Conn) {
	body := `<!DOCTYPE html> <html lang="en"> <head> <meta charset="UTF-8"> <meta name="viewport" content="width=device-width, initial-scale=1.0"> <meta http-equiv="X-UA-Compatible" content="ie=edge"> <title>Document</title> </head> <body> Fall <ul> <li>101 - HTML & CSS - 20</li> <li>102 - JavaScript - 20</li> <li>103 - Jquery - 20</li> </ul> Spring <ul> <li>201 - Golang - 20</li> <li>202 - Golang web dev - 20</li> <li>203 - DevOps - 20</li> </ul> </body> </html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}

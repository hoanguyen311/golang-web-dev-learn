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
	scanner := bufio.NewScanner(conn)
	db := make(map[string]string)

	for scanner.Scan() {
		text := scanner.Text()
		str := strings.Split(text, " ")

		switch str[0] {
		case "GET":
			fmt.Fprintln(conn, "YOUR VALUE", db[str[1]])
		case "SET":
			if len(str) != 3 {
				fmt.Fprintf(conn, "VALUE IS EXPECTED")
				continue
			}
			db[str[1]] = str[2]
		case "DELETE":
			delete(db, str[1])
			fmt.Fprintln(conn, "DELETED ", str[1])
		default:
			fmt.Fprintln(conn, "Invalid command")
		}
	}

	defer conn.Close()
}

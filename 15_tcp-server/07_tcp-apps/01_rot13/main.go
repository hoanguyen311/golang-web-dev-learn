package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := li.Accept()

		if err != nil {
			fmt.Println(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		text := scanner.Text()
		r13 := rotate([]byte(text))

		fmt.Fprintf(conn, "%s - %s\n", text, r13)
	}

	defer conn.Close()
}

func rotate(str []byte) []byte {
	var r13 []byte
	r13 = make([]byte, len(str))

	for i, v := range str {
		if v <= 109 {
			r13[i] = v + 13
		} else {
			r13[i] = v - 13
		}
	}

	return r13
}

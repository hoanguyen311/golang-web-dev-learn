package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	li, err := net.Listen("tcp", ":8080")

	handleError(err)

	for {
		conn, err := li.Accept()

		handleError(err)

		go handle(conn)
	}
}
func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		t := scanner.Text()
		fmt.Printf("%s \n", t)
		fmt.Fprint(conn, "I heard you say ", t, "\n")
	}

	defer conn.Close()
}

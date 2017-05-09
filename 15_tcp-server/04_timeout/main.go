package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))

	if err != nil {
		fmt.Println("Error time out")
	}
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		t := scanner.Text()
		fmt.Printf("%s \n", t)
		fmt.Fprint(conn, "I heard you say ", t, "\n")
	}

	defer conn.Close()
}

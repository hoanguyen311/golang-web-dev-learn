package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Fprint(conn, "I dialed you")
}

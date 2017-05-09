package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello world! you're doing great")
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}

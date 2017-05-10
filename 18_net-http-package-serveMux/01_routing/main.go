package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "Hello")
	case "/dog":
		fmt.Fprintf(w, "Gau gau gau")
	case "/cat":
		fmt.Fprintf(w, "Miao miao")
	default:
		http.NotFound(w, req)
	}
}

func main() {
	var h hotdog

	http.ListenAndServe(":8080", h)
}

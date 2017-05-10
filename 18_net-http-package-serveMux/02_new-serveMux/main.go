package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Gau gau gau")
}

type hotcat int

func (h hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Miao miao")
}

func main() {
	var d hotdog
	var c hotcat

	mux := http.NewServeMux()

	mux.Handle("/dog", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}

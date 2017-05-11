package main

import (
	"fmt"
	"net/http"
)

func d(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Gau gau gau")
}

func c(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Miao miao")
}

func b(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Chip chip chip")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/dog", d)
	mux.HandleFunc("/cat", c)
	mux.Handle("/bird", http.HandlerFunc(b))

	http.ListenAndServe(":8080", mux)
}

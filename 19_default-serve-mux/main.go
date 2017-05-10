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

func main() {
	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)
}

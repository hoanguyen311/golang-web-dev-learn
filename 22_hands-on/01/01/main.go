package main

import (
	"io"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(wr http.ResponseWriter, req *http.Request) {
		io.WriteString(wr, "index")
	})

	http.HandleFunc("/dog", func(wr http.ResponseWriter, req *http.Request) {
		io.WriteString(wr, "gau gau gau")
	})

	http.HandleFunc("/me", func(wr http.ResponseWriter, req *http.Request) {
		io.WriteString(wr, "Harry Nguyen")
	})

	http.ListenAndServe(":8080", nil)
}

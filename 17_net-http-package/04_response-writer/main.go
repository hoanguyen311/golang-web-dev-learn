package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (h hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("HoaNguyenKey", "This is from Hoa Nguyen")

	fmt.Fprint(w, "<h1>Hello how are you?</h1>")
}

func main() {
	var h hotdog

	http.ListenAndServe(":8080", h)
}

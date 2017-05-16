package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/pets", index)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
        <img src="dog.jpg"/>
		<img src="cat.jpg"/>
		<img src="smile.jpg"/>
    `)
}

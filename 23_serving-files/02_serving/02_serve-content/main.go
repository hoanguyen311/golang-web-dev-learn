package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog.jpg", docPic)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
        <img src="/dog.jpg"/>
    `)
}
func docPic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("dog.jpg")
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}
	stat, err := f.Stat()

	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}
	http.ServeContent(w, r, f.Name(), stat.ModTime(), f)
}

package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", foo)
	http.HandleFunc("/dog", dog)
	http.HandleFunc("/resources/dog.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("./dog.gohtml")

	if err != nil {
		log.Fatal(err)
	}

	err = tpl.Execute(w, nil)

	if err != nil {
		log.Fatal(err)
	}
}
func dogPic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/dog.jpg")
}

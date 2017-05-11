package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type viewData struct {
	Title string
	Text  string
}

func main() {

	http.HandleFunc("/", func(wr http.ResponseWriter, req *http.Request) {
		handleError(tpl.ExecuteTemplate(wr, "tpl.gohtml", viewData{"Index", "Index"}))
	})

	http.HandleFunc("/dog", func(wr http.ResponseWriter, req *http.Request) {
		handleError(tpl.ExecuteTemplate(wr, "tpl.gohtml", viewData{"Dog", "Gau gau gau"}))
	})

	http.HandleFunc("/me", func(wr http.ResponseWriter, req *http.Request) {
		handleError(tpl.ExecuteTemplate(wr, "tpl.gohtml", viewData{"Me", "Harry Nguyen"}))
	})

	http.ListenAndServe(":8080", nil)
}

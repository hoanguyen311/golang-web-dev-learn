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

	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog", http.HandlerFunc(dog))
	http.Handle("/me", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}

func index(wr http.ResponseWriter, req *http.Request) {
	handleError(tpl.ExecuteTemplate(wr, "tpl.gohtml", viewData{"Index", "Index"}))
}
func dog(wr http.ResponseWriter, req *http.Request) {
	handleError(tpl.ExecuteTemplate(wr, "tpl.gohtml", viewData{"Dog", "Gau gau gau"}))
}
func me(wr http.ResponseWriter, req *http.Request) {
	handleError(tpl.ExecuteTemplate(wr, "tpl.gohtml", viewData{"Me", "Harry Nguyen"}))
}

package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

var tpl *template.Template

type hotdog int

func (h hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Method        string
		Submissions   url.Values
		URL           *url.URL
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		req.Method,
		req.Form,
		req.URL,
		req.Header,
		req.Host,
		req.ContentLength,
	}

	tpl.Execute(res, data)
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	var h hotdog

	http.ListenAndServe(":8080", h)
}

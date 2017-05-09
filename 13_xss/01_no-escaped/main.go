package main

import (
	"html/template"
	"log"
	"os"
)

type page struct {
	Title   string
	Heading string
	Input   string
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	p := page{
		"text template",
		"Nothing is escaped in text/tempate",
		"<script>alert('ahihi!')</script>",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", p)

	if err != nil {
		log.Fatal(err)
	}
}

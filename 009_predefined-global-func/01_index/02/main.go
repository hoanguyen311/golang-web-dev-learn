package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	skills := []string{"HTML", "CSS", "Golang"}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", struct {
		Name   string
		Skills []string
	}{"Hoa", skills})

	if err != nil {
		log.Fatal(err)
	}
}

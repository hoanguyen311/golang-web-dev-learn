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

type person struct {
	Name string
	Job  string
}

func main() {
	p := person{
		"Hoa",
		"Developer",
	}

	err := tpl.Execute(os.Stdout, p)

	if err != nil {
		log.Fatal(err)
	}
}

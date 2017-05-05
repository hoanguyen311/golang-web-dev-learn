package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type person struct {
	Name string
	Age  int
}
type doubleZero struct {
	person
	LicenseToKill bool
}

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", doubleZero{person{"Hoa", 27}, false})

	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func printTpl(tpl *template.Template, name string) *template.Template {
	err := tpl.ExecuteTemplate(os.Stdout, name, nil)

	if err != nil {
		log.Fatal(err)
	}

	return tpl
}

func main() {
	printTpl(tpl, "one.gohtml")
	printTpl(tpl, "two.gohtml")
	printTpl(tpl, "three.gohtml")
}

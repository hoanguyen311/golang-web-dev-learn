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
	family := map[string]string{
		"Father":   "Hoa",
		"Mother":   "Cuc",
		"Daughter": "Truc",
		"Son":      "Hieu",
	}

	err := tpl.Execute(os.Stdout, family)

	if err != nil {
		log.Fatal(err)
	}
}

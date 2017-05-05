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
	data := []string{"Hoa", "Cuc", "Hieu"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)

	if err != nil {
		log.Fatal(err)
	}
}

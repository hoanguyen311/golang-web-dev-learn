package main

import (
	"log"
	"os"
	"text/template"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	tpl, err := template.ParseFiles("one.gohtml")

	handleError(err)

	tpl, err = tpl.ParseFiles("two.gohtml", "three.gohtml")

	handleError(err)

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)

	handleError(err)

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)

	handleError(err)

	err = tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)

	handleError(err)
}

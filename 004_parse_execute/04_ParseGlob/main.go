package main

import (
	"html/template"
	"log"
	"os"
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	tpl, err := template.ParseGlob("templates/*.gohtml")

	handleError(err)

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)

	handleError(err)

	err = tpl.ExecuteTemplate(os.Stdout, "three.gohtml", nil)

	handleError(err)

}

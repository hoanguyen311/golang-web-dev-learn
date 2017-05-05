package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {

	tpl.Execute(os.Stdout, struct {
		Score1 int
		Score2 int
	}{12, 10})
}

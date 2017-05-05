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

func (p person) DoubleAge() int {
	return p.Age * 2
}

func (p person) TakeArgs(x int) int {
	return x * 2
}

func (p person) Process() int {
	return 7
}
func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", person{"Hoa", 27})

	if err != nil {
		log.Fatal(err)
	}
}

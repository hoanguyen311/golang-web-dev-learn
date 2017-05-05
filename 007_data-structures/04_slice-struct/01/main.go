package main

import (
	"html/template"
	"os"
)

var tpl *template.Template

type person struct {
	Name string
	Job  string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	p1 := person{
		"Hoa",
		"dev",
	}

	p2 := person{
		"Cuc",
		"teacher",
	}

	p3 := person{
		"Hieu",
		"student",
	}

	p4 := person{
		"Truc",
		"student",
	}

	family := []person{p1, p2, p3, p4}

	tpl.Execute(os.Stdout, family)
}

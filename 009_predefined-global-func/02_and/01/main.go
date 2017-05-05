package main

import (
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name    string
	IsAdmin bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	u1 := user{"Hoa", true}
	u2 := user{"Cuc", true}
	u3 := user{"Hieu", false}
	u4 := user{"", true}

	tpl.Execute(os.Stdout, []user{u1, u2, u3, u4})
}

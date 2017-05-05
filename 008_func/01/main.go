package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var fm = template.FuncMap{
	"uc":         strings.ToUpper,
	"firstThree": firstThree,
}
var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}
func firstThree(s string) string {
	s = strings.TrimSpace(s)
	s = s[:3]
	return s
}

type person struct {
	Name string
	Job  string
}
type car struct {
	Manufacturer string
	Doors        int
}

func main() {
	family := []person{
		person{"Hoa", "Developer"},
		person{"Cuc", "Teacher"},
	}

	transport := []car{
		car{"Ford", 5},
		car{"Toyota", 4},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", struct {
		People    []person
		Transport []car
	}{
		family,
		transport,
	})

	if err != nil {
		log.Fatal(err)
	}

}

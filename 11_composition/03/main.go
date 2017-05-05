package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type course struct {
	Number string
	Name   string
	Units  string
}

type seamaster struct {
	Term    string
	Courses []course
}

type year struct {
	Fall, Spring seamaster
}

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	y := year{
		Fall: seamaster{
			Term: "Fall",
			Courses: []course{
				{"101", "HTML & CSS", "20"},
				{"102", "JavaScript", "20"},
				{"103", "Jquery", "20"},
			},
		},
		Spring: seamaster{
			Term: "Spring",
			Courses: []course{
				{"201", "Golang", "20"},
				{"202", "Golang web dev", "20"},
				{"203", "DevOps", "20"},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", y)

	if err != nil {
		log.Fatal(err)
	}
}

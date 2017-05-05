package main

import (
	"html/template"
	"log"
	"os"
)

type dish struct {
	Name  string
	Price float64
}

type meal struct {
	Name   string
	Dishes []dish
}

type menu struct {
	Breakfast meal
	Lunch     meal
	Dinner    meal
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	data := menu{
		meal{
			"Breakfast",
			[]dish{
				{"Banh Mi", 200},
				{"Pho", 200},
				{"Hu Tiu", 3000},
			},
		},
		meal{
			"Lunch",
			[]dish{
				{"Banh Mi", 2000},
				{"Pho", 22000},
				{"Hu Tiu", 22000},
			},
		},
		meal{
			"Dinner",
			[]dish{
				{"Banh Mi", 2000},
				{"Pho", 22000},
				{"Hu Tiu", 22000},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)

	if err != nil {
		log.Fatal(err)
	}
}

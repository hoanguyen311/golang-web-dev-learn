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
type restaurant struct {
	Name string
	Menu menu
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	m1 := menu{
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

	m2 := menu{
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

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", []restaurant{
		{"B.O.C", m1},
		{"Ut ut", m2},
	})

	if err != nil {
		log.Fatal(err)
	}
}

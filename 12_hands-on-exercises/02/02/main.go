package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
}

type region struct {
	Name   string
	Hotels []hotel
}
type regions []region

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	data := regions{
		region{
			"Southern",
			[]hotel{
				{"New World", "Le Lai", "Sai Gon", "70000"},
				{"Majetic", "Le Lai", "Sai Gon", "70000"},
			},
		},
		region{
			"Northern",
			[]hotel{
				{"Majetic", "Hai Ba Trung", "Ha Noi", "50000"},
			},
		},
		region{
			"Central",
			[]hotel{
				{"Drendro Gold", "Tran Phu", "Nha Trang", "212312"},
			},
		},
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatal(err)
	}

}

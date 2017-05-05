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
	Region  string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", []hotel{
		{"Drendro Gold", "Tran Phu", "Nha Trang", "212312", "Central"},
		{"New World", "Le Lai", "Sai Gon", "70000", "Southern"},
		{"Majetic", "Le Lai", "Sai Gon", "70000", "Southern"},
	})
	if err != nil {
		log.Fatal(err)
	}

}

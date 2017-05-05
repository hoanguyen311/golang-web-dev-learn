package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

var fm = template.FuncMap{
	"fdbl": double,
	"fsqr": sqRoot,
	"fsq":  square,
}

func square(x int) float64 {
	return math.Pow(float64(x), 2)
}
func sqRoot(x float64) float64 {
	return math.Sqrt(float64(x))
}

func double(x int) int {
	return x * 2
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 9)

	if err != nil {
		log.Fatal(err)
	}
}

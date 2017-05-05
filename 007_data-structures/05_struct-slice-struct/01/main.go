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

type car struct {
	Manufacture string
	Model       string
	Doors       int
}

// type items struct {
// 	People    []person
// 	Transport []car
// }

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

	c1 := car{
		Manufacture: "Toyota",
		Model:       "Corolla",
		Doors:       4,
	}

	c2 := car{
		Manufacture: "Ford",
		Model:       "F150",
		Doors:       2,
	}

	cars := []car{c1, c2}

	// data := items{
	// 	family,
	// 	cars,
	// }

	data := struct {
		People    []person
		Transport []car
	}{
		family,
		cars,
	}

	tpl.Execute(os.Stdout, data)
}

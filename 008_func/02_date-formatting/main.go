package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

var fm = template.FuncMap{
	"fmtMdy": formatMonthDayYear,
}

func formatMonthDayYear(t time.Time) string {
	return t.Format(time.RFC822)
}
func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())

	if err != nil {
		log.Fatal(err)
	}
}

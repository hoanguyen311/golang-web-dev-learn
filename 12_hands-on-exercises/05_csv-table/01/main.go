package main

import (
	"encoding/csv"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var tpl *template.Template

type record struct {
	Date time.Time
	Open float64
}

type viewData struct {
	Records []record
	Heading []string
}

func init() {
	tpl = template.Must(tpl.ParseFiles("tpl.gohtml"))
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	data := readFile("table.csv")

	err := tpl.ExecuteTemplate(res, "tpl.gohtml", data)

	if err != nil {
		log.Fatal(err)
	}
}

func readFile(src string) viewData {
	file, _ := os.Open(src)
	r := csv.NewReader(file)

	defer file.Close()

	rows, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}
	records := make([]record, 0, len(rows))

	for _, row := range rows[1:] {
		date, _ := time.Parse("2006-01-02", row[0])
		open, _ := strconv.ParseFloat(row[1], 64)

		records = append(records, record{
			Date: date,
			Open: open,
		})
	}

	return viewData{
		Records: records,
		Heading: rows[0],
	}
}

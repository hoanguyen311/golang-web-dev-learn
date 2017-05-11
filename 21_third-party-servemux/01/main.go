package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	mux := httprouter.New()

	mux.GET("/apply", apply)
	mux.POST("/apply", processApply)
	mux.GET("/contact", contact)
	mux.GET("/blog/:name", blog)

	http.ListenAndServe(":8080", mux)
}

func apply(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func processApply(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()

	data := struct {
		Name string
		Age  string
	}{
		r.PostForm.Get("name"),
		r.PostForm.Get("age"),
	}

	err := tpl.ExecuteTemplate(w, "processApply.gohtml", data)

	if err != nil {
		log.Fatal(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)

	if err != nil {
		log.Fatal(err)
	}
}

func blog(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := tpl.ExecuteTemplate(w, "blog.gohtml", struct {
		Name string
	}{p.ByName("name")})

	if err != nil {
		log.Fatal(err)
	}
}

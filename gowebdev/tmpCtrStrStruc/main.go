package main

import (
	"html/template"
	"net/http"
)

//  {{/* a comment */}}	Defines a comment
/*
{{.}}	Renders the root element
{{.Name}}	Renders the “Name”-field in a nested element
{{if .Done}} {{else}} {{end}}	Defines an if/else-Statement
{{range .List}} {{.}} {{end}}	Loops over all “List” field and renders each using {{.}}
*/

type prodSpec struct {
	Size   string
	Weight float32
	Descr  string
}

type product struct {
	ProdID int
	// Name   string
	Name  string
	Cost  float64
	Specs prodSpec
}

var tpl *template.Template
var prod1 product

func main() {
	prod1 = product{
		ProdID: 15,
		Name:   "Wicked Cool Phone",
		Cost:   899,
		Specs: prodSpec{
			Size:   "150 x 70 x 7 mm",
			Weight: 65,
			Descr:  "Over priced shiny thing designed to shatter on impact",
		},
	}

	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/productinfo", productInfoHandler)
	http.ListenAndServe(":8080", nil)
}

func productInfoHandler(w http.ResponseWriter, r *http.Request) {
	// func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	tpl.ExecuteTemplate(w, "productinfo2.html", prod1)
}

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

// GroceryList data type for export
type GroceryList []string

var tpl *template.Template
var g GroceryList

func main() {
	g = GroceryList{"milk", "eggs", "green beans", "cheese", "flour", "sugar", "broccoli"}
	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/list", listHandler)
	http.ListenAndServe(":8080", nil)
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	// func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	tpl.ExecuteTemplate(w, "groceries.html", g)
}

package main

import (
	"fmt"
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

var tpl *template.Template
var name = "John"

func main() {
	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/welcome", welcomeHandler)
	http.ListenAndServe(":8080", nil)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("indexHandler running")
	// func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	tpl.ExecuteTemplate(w, "welcome.html", name)
}

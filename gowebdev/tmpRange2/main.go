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

type task struct {
	Name string
	Done bool
}

var tpl *template.Template

// Todo list to be exported
var Todo []task

func main() {
	Todo = []task{{"give dog a bath", true}, {"mow the lawn", false}, {"pickup groceries", false}, {"take out trash", false}, {"paint kitchen", false}, {"feed dog", false}, {"water plants", false}}
	// Todo = []task{{"give dog a bath", true}, {"mow the lawn", true}, {"pickup groceries", true}, {"take out trash", true}, {"paint living room", false}, {"feed dog", true}, {"water plants", true}}
	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/todo", todoHandler)
	http.ListenAndServe(":8080", nil)
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	// func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	tpl.ExecuteTemplate(w, "todolist.html", Todo)
}

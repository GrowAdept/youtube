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

var tpl *template.Template

// User first letter must be capitalized to be exported
type User struct {
	Name     string
	Language string
	Member   bool
}

// U struct for export
var u User

func main() {
	u = User{"Bob Smith", "English", false}
	// u = User{"Juan Hernández", "Spanish", true}
	// u = User{"Zhang Wei", "Mandarin", true}
	// u = User{"007", "?", true}

	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/welcome", welcomeHandler)
	http.ListenAndServe(":8080", nil)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	// func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	tpl.ExecuteTemplate(w, "membership.html", u)
}

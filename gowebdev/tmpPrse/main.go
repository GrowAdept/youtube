package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	// func ParseGlob(pattern string) (*Template, error)
	// tpl, _ = template.ParseGlob("templates/*.html")
	// func (t *Template) ParseFiles(filenames ...string) (*Template, error)
	tpl, _ = tpl.ParseGlob("templates/*.html")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)
	http.HandleFunc("/login", loginHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// func (t *Template) ExecuteTemplate(wr io.Writer, name string, data interface{}) error
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.html", nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "contact.html", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "login.html", nil)
}

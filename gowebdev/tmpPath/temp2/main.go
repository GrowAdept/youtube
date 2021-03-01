package main

import (
	"html/template"
	"net/http"
)

/*
-tmpPath
  -index4.html
  -temp2
	-main.go  <----- starting app here
	-index1.html
	-data1
	  -index2.html
	  -data2
	    -index3.html
*/

var tpl *template.Template

func main() {
	// func ParseFiles(filenames ...string) (*Template, error)
	tpl, _ = template.ParseFiles("index1.html")
	// tpl, _ = template.ParseFiles("data1/index2.html")
	// tpl, _ = template.ParseFiles("data1/data2/index3.html")
	// tpl, _ = template.ParseFiles("../index4.html")
	// func (t *Template) ParseFiles(filenames ...string) (*Template, error)
	// tpl, _ = tpl.ParseFiles("../index4.html")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// func (t *Template) Execute(wr io.Writer, data interface{}) error
	tpl.Execute(w, nil)
}

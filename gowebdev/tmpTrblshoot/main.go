package main

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

// first letter of var type being uppercase or lowercase doesn't affect usage in templates
type price struct {
	// first letter of struct field name does affect usage in templates
	UppercaseField float64
	lowercaseField float64
}

var funcMap = template.FuncMap{
	"CanCashPr": CanCashPr,
	"Upper":     strings.ToUpper,
}

var tpl *template.Template

// first letter of var name being uppercase or lowercase doesn't affect usage in templates
var p price

func CanCashPr(p float64) (str string, err error) {
	err = errors.New("warning!")
	remainder := int(p*100) % 5
	quotiant := int(p*100) / 5
	if remainder < 3 {
		pr := float64(quotiant*5) / 100
		s := fmt.Sprintf("%.2f", pr)
		return s, err
	}
	pr := (float64(quotiant*5) + 5) / 100
	s := fmt.Sprintf("%.2f", pr)
	return s, err
}

func main() {
	// New() allocates a new HTML template with the given name
	// no file is parsed yet
	// func New(name string) *Template
	tpl = template.New("index.html")

	// Funcs adds the elements of the argument map to the template's function map.
	// func (t *Template) Funcs(funcMap FuncMap) *Template
	tpl = tpl.Funcs(funcMap)

	// parse our files
	// by not checking for parsing errors get a confusing stack trace
	// tpl, _ = tpl.ParseFiles("templates/index.html")
	// func Must(t *Template, err error) *Template
	// func (t *Template) ParseFiles(filenames ...string) (*Template, error)
	tpl = template.Must(tpl.ParseFiles("templates/index.html"))

	/*
		var err error
		tpl, err = tpl.ParseFiles("templates/index.html")
		if err != nil {
			panic(err)
		}
	*/

	p.UppercaseField = 3.3333333
	p.lowercaseField = 3.3333333

	// fmt.Println("tpl.Tree", *tpl.Tree)
	http.HandleFunc("/index1", indexHandler1)
	http.HandleFunc("/index2", indexHandler2)
	http.HandleFunc("/error", errorHandler)
	http.ListenAndServe("localhost:3000", nil)
}

func indexHandler1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--------indexHandler1 running--------")
	// executes template and writes template with var p
	tpl.ExecuteTemplate(w, "index.html", p)
}

// indexHandler2 seperates the executing of the template and writing to the http.Responswiter
// in order to better check for errors while execting the template
func indexHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("--------indexHandler2 running--------")
	// io.Writer is an interaface with the method Write(p []byte) (n int, err error)
	var buff bytes.Buffer
	// func (t *Template) Execute(wr io.Writer, data interface{}) error
	err := tpl.Execute(&buff, p)
	if err != nil {
		fmt.Println("err:", err)
		http.Redirect(w, r, "/error", http.StatusTemporaryRedirect)
		return
	}
	buff.WriteTo(w)
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "there was an error parsing page")
}

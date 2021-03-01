package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

// Price of item
type Price float64

// CanCashPr converts Canadian price to cash price (no 1 cent coins)
func (p Price) CanCashPr() string {
	remainder := int(p*100) % 5
	quotiant := int(p*100) / 5
	if remainder < 3 {
		pr := float64(quotiant*5) / 100
		s := fmt.Sprintf("%.2f", pr)
		return s
	}
	pr := (float64(quotiant*5) + 5) / 100
	s := fmt.Sprintf("%.2f", pr)
	return s
}

var p Price

func main() {
	p = 3.91
	tpl, _ = tpl.ParseFiles("index.html")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.html", p)
}

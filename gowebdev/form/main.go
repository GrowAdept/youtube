package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Sub struct {
	Username string
	Data     string
}

var tpl *template.Template

func main() {
	tpl, _ = tpl.ParseGlob("templates/*.html")
	http.HandleFunc("/getform", getFormHandler)
	http.HandleFunc("/processget", processGetHandler)
	http.HandleFunc("/postform", postFormHandler)
	http.HandleFunc("/processpost", processPostHandler)
	http.ListenAndServe(":8080", nil)
}

func getFormHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "getform.html", nil)
}

func processGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("processGetHandler running")
	var s Sub
	s.Username = r.FormValue("usernameName")
	s.Data = r.FormValue("dataName")
	fmt.Println("Username:", s.Username, "Sensitive Data:", s.Data)
	tpl.ExecuteTemplate(w, "thanks.html", s)
}

func postFormHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "postform.html", nil)
}

func processPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("processPostHandler running")
	var s Sub
	s.Username = r.FormValue("usernameName")
	s.Data = r.FormValue("dataName")
	fmt.Println("Username:", s.Username, "Sensitive Data:", s.Data)
	tpl.ExecuteTemplate(w, "thanks.html", s)
}

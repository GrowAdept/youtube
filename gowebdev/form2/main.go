package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Sub type holds user submissions
type Sub struct {
	Username string
	Num      int
	MyFloat  float64
	Updates  bool
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
	s.Username = r.FormValue("username")
	// cannot use r.FormValue("numberName") (type string) as type int in assignment
	// s.Num = r.FormValue("numberName")
	num := r.FormValue("numberName")
	// ASCII to int
	// func Atoi(s string) (int, error)
	s.Num, _ = strconv.Atoi(num)
	s.Num = s.Num * 2
	var err error
	/*
		// invalid syntax cannot parse float64
		s.Num, err = strconv.Atoi(r.FormValue("myFltName"))
		fmt.Println("error:", err)
	*/
	// func ParseFloat(s string, bitSize int) (float64, error)
	s.MyFloat, err = strconv.ParseFloat(r.FormValue("myFltName"), 64)
	if err != nil {
		log.Fatal("error parsing float64")
	}
	if r.FormValue("upName") == "true" {
		s.Updates = true
	} else if r.FormValue("upName") == "false" {
		s.Updates = false
	}
	tpl.ExecuteTemplate(w, "thanks.html", s)
}

func postFormHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "postform.html", nil)
}

func processPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("processPostHandler running")

	var s Sub
	s.Username = r.FormValue("username")
	// cannot use r.FormValue("numberName") (type string) as type int in assignment
	// s.Num = r.FormValue("numberName")
	num := r.FormValue("numberName")
	// ASCII to int
	// func Atoi(s string) (int, error)
	s.Num, _ = strconv.Atoi(num)
	s.Num = s.Num * 2
	var err error
	// func ParseFloat(s string, bitSize int) (float64, error)
	s.Num, _ = strconv.Atoi(r.FormValue("myFltName"))
	//s.MyFloat, err = strconv.ParseFloat(r.FormValue("myFltName"), 64)
	if err != nil {
		log.Fatal("error parsing float64")
	}
	if r.FormValue("upName") == "true" {
		s.Updates = true
	} else if r.FormValue("upName") == "false" {
		s.Updates = false
	}
	tpl.ExecuteTemplate(w, "thanks.html", s)
}

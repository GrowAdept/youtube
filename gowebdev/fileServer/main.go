package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func main() {
	tpl, _ = tpl.ParseGlob("templates/*.html")
	// create our new var myDir at type http.Dir
	myDir := http.Dir("/workspace/goworkspace/src/gowebdev/fileServer/public")
	fmt.Printf("myDir type: %T", myDir)
	// func FileServer(root FileSystem) Handler
	myHandler := http.FileServer(myDir)
	http.Handle("/", myHandler)
	// using absolute path
	// http.Handle("/", http.FileServer(http.Dir("/workspace/goworkspace/src/gowebdev/fileServer/public")))
	// using relative path
	// http.Handle("/", http.FileServer(http.Dir("./public")))
	// does not work, will look at ./public/public
	// http.Handle("/public", http.FileServer(http.Dir("./public")))
	// use http.StringPrefix to alter request before FileServer sees it
	// http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	// http.Handle("/public/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "hello.html", nil)
}

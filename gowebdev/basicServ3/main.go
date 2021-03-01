package main

import (
	"fmt"
	"net/http"
)

func main() {
	// func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	http.HandleFunc("/hello", helloHandleFunc)
	http.HandleFunc("/about", aboutHandleFunc)
	// func ListenAndServe(addr string, handler Handler) error
	http.ListenAndServe(":8080", nil)
}

func helloHandleFunc(w http.ResponseWriter, r *http.Request) {
	// func Fprint(w io.Writer, a ...interface{}) (n int, err error)
	// fmt.Fprint(w, "Hello, World!")
	fmt.Fprintf(w, "r.URL.Path:, %s", r.URL.Path)
}

func aboutHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to our Gopher powered website.")
}

// basic web server
package main

import (
    "fmt"
    "log"
    "net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the home page :)")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the About Page.  Not much to see here, just a basic web app.")
}

func main() {
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/about", aboutHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
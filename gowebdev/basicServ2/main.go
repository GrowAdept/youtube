package main

import (
	"fmt"
	"net/http"
)

func main() {
	// registers the handler function for the given pattern in the DefaultServeMux.
	http.HandleFunc("/hello", helloHandleFunc)
	// ListenAndServe listens on the TCP network address addr and then calls Serve
	// with handler to handle requests on incoming connections
	http.ListenAndServe("localhost:8080", nil) // entering nill implicitly uses DefaultServeMux
	// ServeMux is an HTTP request multiplexer. It matches the URL of each incoming
	// request against a list of registered patterns and calls the handler for the
	// pattern that most closely matches the URL.
}

// handler function that responds to client http requests
func helloHandleFunc(w http.ResponseWriter, r *http.Request) {
	// formats using the default formats for its operands and writes to w
	fmt.Fprint(w, "Hello, World!")
}

package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type gopher struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var gophers = []gopher{
	{"1", "Ken", "Thompson"},
	{"2", "Robert", "Griesemer"},
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/gopher", getGopher).Methods("GET")
	router.HandleFunc("/gopher", createGopher).Methods("POST")
	http.Handle("/", router)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

/*
w.Header.Set() sets the content type to application/json
if w.WriteHeader is no called explicily, implicit http.StatusOK is used with w.Write()
jsonGohper, err := json.Marshal(gopher) converts the gopher struct to JSON
if err != nil {} sends the 400 error code if the marshaling fails and ends the func
w.Writer()  writes the data to the connection as part of an HTTP reply
*/
func getGopher(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonGohper, err := json.Marshal(gophers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(jsonGohper)
}

/*
A Decoder reads and decodes JSON values from an input stream.
NewDecoder returns a new decoder that reads from r.
The decoder introduces its own buffering and may read data from r beyond the JSON values requested.
Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by v.
handle any errors create decoder, setting HTTP status code
since err from defer r.Body.Close() is not nill, closing r.Body must be explicitly closed
json.Marshal(gophers) marshals gophers slice so w.Write() can recieve it
w.Header().Set() sets the headers to application/json
w.Write() writes our JSON to the http.ResponseWriter
*/
func createGopher(w http.ResponseWriter, r *http.Request) {
	var newGohper gopher
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newGohper)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	gophers = append(gophers, newGohper)
	GophersJSON, _ := json.Marshal(gophers)
	w.Header().Set("Content-Type", "application/json")
	w.Write(GophersJSON)
}

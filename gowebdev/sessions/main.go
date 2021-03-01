// simplified example and not intended for production use
package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

var tpl *template.Template

// func NewCookieStore(keyPairs ...[]byte) *CookieStore
var store = sessions.NewCookieStore([]byte("super-secret-password"))

func main() {
	tpl, _ = template.ParseGlob("templates/*.html")
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/about", aboutHandler)
	http.ListenAndServe("localhost:8080", context.ClearHandler(http.DefaultServeMux))
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	// Get a session. Get() always returns a session, even if empty.
	// func (s *FilesystemStore) Get(r *http.Request, name string) (*Session, error)
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options = &sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,
		// MaxAge:   5,
		HttpOnly: true,
	}
	r.ParseForm()
	name := r.FormValue("name")
	if name != "" {
		// Set name session value.
		session.Values["name"] = name
	}
	fmt.Println("session:", session)
	// Save it before we write to the response/return from the handler.
	// func (s *Session) Save(r *http.Request, w http.ResponseWriter) error
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.ExecuteTemplate(w, "create.html", name)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.html", nil)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	// func (s *FilesystemStore) Get(r *http.Request, name string) (*Session, error)
	session, _ := store.Get(r, "session-name")
	session.Options.MaxAge = -1
	fmt.Println("session:", session)
	// call Save before writing to the response or returning from the handler
	// func (s *Session) Save(r *http.Request, w http.ResponseWriter) error
	session.Save(r, w)
	tpl.ExecuteTemplate(w, "delete.html", nil)
}

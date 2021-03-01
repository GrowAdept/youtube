package main

import (
	"fmt"
	"html/template"
	"net/http"
)

/*
type Cookie struct {
    Name  string
    Value string

    Path       string    - indicates a URL path that must exist in req URL to send Cookie Header
    Domain     string    - specifies which hosts are allowed to recieve the cookie
    Expires    time.Time - deletes at specified date
    RawExpires string    // for reading cookies only

    // MaxAge=0 means no 'Max-Age' attribute specified.
    // MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
    // MaxAge>0 means Max-Age attribute present and given in seconds
    MaxAge   int		- deletes cookie after specified amount of time in seconds
    Secure   bool		- sent to server on encrypted req, never unsecured (HTTP)
    HttpOnly bool		- not accessible by JavaScript, only sent to server
    SameSite SameSite 	- servers require that a cookie shouldn't be sent with cross-orign req
    Raw      string
    Unparsed []string   - raw text of unparsed attribute-value pairs
}
*/

var tpl *template.Template

func main() {
	tpl, _ = template.ParseGlob("templates/*.html")
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe("localhost:8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	//  looks for cookie on our machine
	//  func (r *Request) Cookie(name string) (*Cookie, error)
	cookie, err := r.Cookie("2nd-cookie")
	fmt.Println("cookie:", cookie, "err:", err)
	if err != nil {
		fmt.Println("cookie was not found")
		cookie = &http.Cookie{
			Name:     "2nd-cookie",
			Value:    "my second cookie value",
			HttpOnly: true,
		}
		// func SetCookie(w ResponseWriter, cookie *Cookie)
		http.SetCookie(w, cookie)
	}
	tpl.ExecuteTemplate(w, "index.html", nil)
}

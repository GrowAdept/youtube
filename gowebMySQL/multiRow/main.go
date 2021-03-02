package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Product data type for export
type Product struct {
	ID          int
	Name        string
	Price       float32
	Description string
}

var tpl *template.Template

var db *sql.DB

func main() {
	tpl, _ = template.ParseGlob("templates/*.html")
	var err error
	// Never use _, := db.Open(), release resources with db.Close
	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	http.HandleFunc("/productsearch", productSearchHandler)
	http.HandleFunc("/productsearch2", productSearchHandler2)
	http.HandleFunc("/", homePageHandler)
	http.ListenAndServe("localhost:8080", nil)
}

func productSearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "productsearch.html", nil)
		return
	}
	r.ParseForm()
	// func (r *Request) FormValue(key string) string
	min := r.FormValue("minPriceName")
	max := r.FormValue("maxPriceName")
	stmt := "SELECT * FROM products WHERE price >= ? && price <= ?;"
	// func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
	rows, err := db.Query(stmt, min, max)
	if err != nil {
		panic(err)
	}
	// type sql.Row does not have a .Close() Method but sql.Rows does and must be run
	defer rows.Close()
	var products []Product
	// func (rs *Rows) Next() bool
	for rows.Next() {
		var p Product
		// func (rs *Rows) Scan(dest ...interface{}) error
		err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description)
		if err != nil {
			panic(err)
		}
		// func append(slice []Type, elems ...Type) []Type
		products = append(products, p)
	}
	tpl.ExecuteTemplate(w, "productsearch.html", products)
}

func productSearchHandler2(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "productsearch2.html", nil)
		return
	}
	r.ParseForm()
	// func (r *Request) FormValue(key string) string
	min := r.FormValue("minPriceName")
	max := r.FormValue("maxPriceName")
	//  func (db *DB) Prepare(query string) (*Stmt, error)
	stmt, err := db.Prepare("SELECT * FROM products WHERE price >= ? && price <= ?;")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	//  func (s *Stmt) Query(args ...interface{}) (*Rows, error)
	rows, err := stmt.Query(min, max)
	var products []Product
	// func (rs *Rows) Next() bool
	for rows.Next() {
		var p Product
		// func (rs *Rows) Scan(dest ...interface{}) error
		err = rows.Scan(&p.ID, &p.Name, &p.Price, &p.Description)
		if err != nil {
			panic(err)
		}
		// func append(slice []Type, elems ...Type) []Type
		products = append(products, p)
	}
	tpl.ExecuteTemplate(w, "productsearch2.html", products)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

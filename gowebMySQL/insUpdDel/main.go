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
	// Never use _, := db.Open(), resources need to be released with db.Close
	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	http.HandleFunc("/insert", insertHandler)
	http.HandleFunc("/browse", browseHandler)
	http.HandleFunc("/update/", updateHandler)
	http.HandleFunc("/updateresult/", updateResultHandler)
	http.HandleFunc("/delete/", deleteHandler)
	http.HandleFunc("/", homePageHandler)
	http.ListenAndServe("localhost:8080", nil)
}

func browseHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****browseHandler running*****")
	stmt := "SELECT * FROM products"
	// func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
	rows, err := db.Query(stmt)
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
	tpl.ExecuteTemplate(w, "select.html", products)
}

func insertHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****insertHandler running*****")
	if r.Method == "GET" {
		tpl.ExecuteTemplate(w, "insert.html", nil)
		return
	}
	r.ParseForm()
	// func (r *Request) FormValue(key string) string
	name := r.FormValue("nameName")
	price := r.FormValue("priceName")
	descr := r.FormValue("descrName")
	var err error
	if name == "" || price == "" || descr == "" {
		fmt.Println("Error inserting row:", err)
		tpl.ExecuteTemplate(w, "insert.html", "Error inserting data, please check all fields.")
		return
	}
	var ins *sql.Stmt
	// don't use _, err := db.Query()
	// func (db *DB) Prepare(query string) (*Stmt, error)
	ins, err = db.Prepare("INSERT INTO `testdb`.`products` (`name`, `price`, `description`) VALUES (?, ?, ?);")
	if err != nil {
		panic(err)
	}
	defer ins.Close()
	// func (s *Stmt) Exec(args ...interface{}) (Result, error)
	res, err := ins.Exec(name, price, descr)

	// check rows affectect???????
	rowsAffec, _ := res.RowsAffected()
	if err != nil || rowsAffec != 1 {
		fmt.Println("Error inserting row:", err)
		tpl.ExecuteTemplate(w, "insert.html", "Error inserting data, please check all fields.")
		return
	}
	lastInserted, _ := res.LastInsertId()
	rowsAffected, _ := res.RowsAffected()
	fmt.Println("ID of last row inserted:", lastInserted)
	fmt.Println("number of rows affected:", rowsAffected)
	tpl.ExecuteTemplate(w, "insert.html", "Product Successfully Inserted")
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****updateHandler running*****")
	r.ParseForm()
	id := r.FormValue("idproducts")
	row := db.QueryRow("SELECT * FROM testdb.products WHERE idproducts = ?;", id)
	var p Product
	// func (r *Row) Scan(dest ...interface{}) error
	err := row.Scan(&p.ID, &p.Name, &p.Price, &p.Description)
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/browse", 307)
		return
	}
	tpl.ExecuteTemplate(w, "update.html", p)
}

func updateResultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****updateResultHandler running*****")
	r.ParseForm()
	id := r.FormValue("idproducts")
	name := r.FormValue("nameName")
	price := r.FormValue("priceName")
	description := r.FormValue("descrName")
	upStmt := "UPDATE `testdb`.`products` SET `name` = ?, `price` = ?, `description` = ? WHERE (`idproducts` = ?);"
	// func (db *DB) Prepare(query string) (*Stmt, error)
	stmt, err := db.Prepare(upStmt)
	if err != nil {
		fmt.Println("error preparing stmt")
		panic(err)
	}
	fmt.Println("db.Prepare err:", err)
	fmt.Println("db.Prepare stmt:", stmt)
	defer stmt.Close()
	var res sql.Result
	// func (s *Stmt) Exec(args ...interface{}) (Result, error)
	res, err = stmt.Exec(name, price, description, id)
	rowsAff, _ := res.RowsAffected()
	if err != nil || rowsAff != 1 {
		fmt.Println(err)
		tpl.ExecuteTemplate(w, "result.html", "There was a problem updating the product")
		return
	}
	tpl.ExecuteTemplate(w, "result.html", "Product was Successfully Updated")
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****deleteHandler running*****")
	r.ParseForm()
	id := r.FormValue("idproducts")
	//  func (db *DB) Prepare(query string) (*Stmt, error)
	del, err := db.Prepare("DELETE FROM `testdb`.`products` WHERE (`idproducts` = ?);")
	if err != nil {
		panic(err)
	}
	defer del.Close()
	var res sql.Result
	res, err = del.Exec(id)
	rowsAff, _ := res.RowsAffected()
	fmt.Println("rowsAff:", rowsAff)

	if err != nil || rowsAff != 1 {
		fmt.Fprint(w, "Error deleting product")
		return
	}
	/*
		if err != nil {
			fmt.Fprint(w, "Error deleting product")
			return
		}
	*/
	fmt.Println("err:", err)
	tpl.ExecuteTemplate(w, "result.html", "Product was Successfully Deleted")
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/browse", 307)
}

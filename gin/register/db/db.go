package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:super-secret-password@tcp(localhost:3306)/gin_db")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}

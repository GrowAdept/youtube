package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// successful db connection or panic
func init() {
	dsn := MySQLusername + ":" + MySQLpassword + "@tcp(" + MySQLaddress + ":" + MySQLport + ")/" + dbName
	// sql.Open will validate arguements but does not create connection
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	// sql.Ping verifies a connection to the database
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
}

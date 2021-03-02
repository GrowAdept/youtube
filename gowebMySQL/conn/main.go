package main

import (
	"database/sql"
	"fmt"
	f "fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Open may just validate its arguments without creating a connection to the database.
	// To verify that the data source name is valid, call Ping.
	// func Open(driverName, dataSourceName string) (*DB, error)
	// dataSourceName: username:password@protocol(address)/dbname?param=value
	//db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/testdb")

	fmt.Println(os.Getenv("MYSQL_PASSWORD"))
	pswd := os.Getenv("MYSQL_PASSWORD")
	db, err := sql.Open("mysql", "root:"+pswd+"@tcp(localhost:3306)/testdb")

	if err != nil {
		f.Println("error validating sql.Open arguments")
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		f.Println("error verifying connection with db.Ping")
		panic(err.Error())
	}

	// func (db *DB) Query(query string, args ...interface{}) (*Rows, error)
	insert, err := db.Query("INSERT INTO `testdb`.`students` (`id`, `firstname`, `lastname`) VALUES ('3', 'Carl', 'Jones');")
	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	f.Println("Successful Connection to Database!")
}

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	var err error
	db, err = sql.Open("mysql", "root:password@tcp(localhost:3306)/testdb")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// func transSavToCheck(userID int, amount float32)
	transSavToCheck(2, 100)
}

func transSavToCheck(userID int, amount float32) {
	// func (db *DB) Begin() (*Tx, error)
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	var checkBalance float32
	var savBalance float32
	// func (tx *Tx) QueryRow(query string, args ...interface{}) *Row
	row := tx.QueryRow("SELECT balance FROM savings WHERE user_id = ?", userID)
	err = row.Scan(&savBalance)
	if err != nil {
		// func (tx *Tx) Rollback() error
		_ = tx.Rollback()
		fmt.Println("err:", err)
		return
	}
	// func (tx *Tx) QueryRow(query string, args ...interface{}) *Row
	row = tx.QueryRow("SELECT balance FROM checking WHERE user_id = ?", userID)
	err = row.Scan(&checkBalance)
	if err != nil {
		// func (tx *Tx) Rollback() error
		_ = tx.Rollback()
		fmt.Println("err:", err)
		return
	}
	savBalance = savBalance - amount
	checkBalance = checkBalance + amount
	fmt.Println("attempting to set checking:", checkBalance, "savings:", savBalance)
	var result sql.Result
	// func (tx *Tx) Exec(query string, args ...interface{}) (Result, error)
	result, execErr := tx.Exec(`UPDATE savings SET balance = ? WHERE user_id = ?`, savBalance, userID)
	rowsAffected, _ := result.RowsAffected()
	fmt.Println("update savings execErr:", execErr, "rowsAffected:", rowsAffected)
	if execErr != nil || rowsAffected != 1 {
		//  func (tx *Tx) Rollback() error
		_ = tx.Rollback()
		return
	}
	// func (tx *Tx) Exec(query string, args ...interface{}) (Result, error)
	result, execErr = tx.Exec(`UPDATE checking SET balance = ? WHERE user_id = ?`, checkBalance, userID)
	rowsAffected, _ = result.RowsAffected()
	fmt.Println("update savings execErr:", execErr, "rowsAffected:", rowsAffected)
	if execErr != nil || rowsAffected != 1 {
		//  func (tx *Tx) Rollback() error
		_ = tx.Rollback()
		return
	}
	// func (tx *Tx) Commit() error
	if err := tx.Commit(); err != nil {
		fmt.Println(err)
	}
}

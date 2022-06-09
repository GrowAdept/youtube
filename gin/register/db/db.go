package db

import (
	"database/sql"
	"fmt"
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

func (u *User) GetUserByUsername() error {
	stmt := "SELECT * FROM users WHERE username = ?"
	row := db.QueryRow(stmt, u.Username)
	err := row.Scan(&u.ID, &u.Username, &u.Email, &u.pswdHash, &u.CreatedAt, &u.Active, &u.verHash, &u.timeout)
	if err != nil {
		fmt.Println("getUser() error selecting User, err:", err)
		return err
	}
	return nil
}

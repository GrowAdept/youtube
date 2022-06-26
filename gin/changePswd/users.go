package main

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"unicode"

	passwordvalidator "github.com/wagslane/go-password-validator"
	"golang.org/x/crypto/bcrypt"
)

func (u *User) getUserByUsername() error {
	stmt := "SELECT * FROM users WHERE username = ?"
	row := db.QueryRow(stmt, u.Username)
	err := row.Scan(&u.ID, &u.Username, &u.Email, &u.pswdHash, &u.CreatedAt, &u.Active, &u.verHash, &u.timeout)
	if err != nil {
		fmt.Println("getUserByUsername() error selecting User, err:", err)
		return err
	}
	return nil
}

func (u *User) getUserByID() error {
	stmt := "SELECT * FROM users WHERE id = ?"
	row := db.QueryRow(stmt, u.ID)
	err := row.Scan(&u.ID, &u.Username, &u.Email, &u.pswdHash, &u.CreatedAt, &u.Active, &u.verHash, &u.timeout)
	if err != nil {
		fmt.Println("getUserByID() error selecting User, err:", err)
		return err
	}
	return nil
}

func (u *User) getUserByEmail() error {
	stmt := "SELECT * FROM users WHERE email = ?"
	row := db.QueryRow(stmt, u.Email)
	err := row.Scan(&u.ID, &u.Username, &u.Email, &u.pswdHash, &u.CreatedAt, &u.Active, &u.verHash, &u.timeout)
	if err != nil {
		fmt.Println("getUserByEmail() error selecting User, err:", err)
		return err
	}
	return nil
}

// validateUsername checks username only has alphanumeric characters
// and if sufficient length, errors are safe to share with user
func (u *User) validateUsername() error {
	// check username for only alphaNumeric characters
	for _, char := range u.Username {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			return errors.New("only alphanumeric characters allowed for username")
		}
	}
	// check username length
	if 5 <= len(u.Username) && len(u.Username) <= 50 {
		return nil
	}
	return errors.New("username length must be greater than 4 and less than 51 characters")
}

// validatePassword checks level of entropy of password
func (u *User) validatePassword() error {
	// if the password has enough entropy, err is nil
	// otherwise, a formatted error message is provided explaining
	// how to increase the strength of the password
	// (safe to show to the client)
	err := passwordvalidator.Validate(u.password, minEntropyBits)
	return err
}

func (u *User) UsernameExists() (exists bool) {
	exists = true
	stmt := "SELECT id FROM users WHERE username = ?"
	row := db.QueryRow(stmt, u.Username)
	var uID string
	err := row.Scan(&uID)
	if err == sql.ErrNoRows {
		return false
	}
	return exists
}

// EmailExists checks if email address exists in db
func (u *User) EmailExists() (exists bool) {
	exists = true
	stmt := "SELECT id FROM users WHERE email = ?"
	row := db.QueryRow(stmt, u.Email)
	var uID string
	err := row.Scan(&uID)
	if err == sql.ErrNoRows {
		return false
	}
	return exists
}

// New creates a new User in the database, is_active set to 0 (false) until email is verfied
func (u *User) New() error {
	// create hash from password
	var hash []byte
	hash, err := bcrypt.GenerateFromPassword([]byte(u.password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt err:", err)
		return err
	}
	// create user value
	CreatedAt := time.Now().Local()
	// return time to the nanosecond (1 billionth of a sec)
	rand.Seed(time.Now().UnixNano())
	// create random code for email
	// Go rune data type represent Unicode characters
	timeout := time.Now().Local().AddDate(0, 0, 2)
	var emailVerPassword string
	emailVerPassword, u.verHash, err = u.NewEmailVerPswd()
	if err != nil {
		return nil
	}
	// save user
	var tx *sql.Tx
	tx, err = db.Begin()
	if err != nil {
		fmt.Println("failed to begin transaction, err:", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return err
	}
	defer tx.Rollback()
	var insertStmt *sql.Stmt
	insertStmt, err = tx.Prepare("INSERT INTO users (username, email, pswd_hash, created_at, is_active, ver_hash, timeout) VALUES (?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
			return err
		}
		defer insertStmt.Close()
	}
	var result sql.Result
	// check if user exists already
	result, err = insertStmt.Exec(u.Username, u.Email, hash, CreatedAt, 0, u.verHash, timeout)
	aff, _ := result.RowsAffected()
	if aff == 0 {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
			return err
		}
		return err
	}
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
			return err
		}
		return err
	}
	// send email
	subject := "Email Verificaion"
	HTMLbody :=
		`<html>
			<h1>Click Link to Veify Email</h1>
			<a href="` + domName + `/emailver/` + u.Username + `/` + emailVerPassword + `">click to verify email</a>
		</html>`
	err = u.SendEmail(subject, HTMLbody)
	if err != nil {
		fmt.Println("issue sending verification email")
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return err
	}
	if commitErr := tx.Commit(); commitErr != nil {
		fmt.Println("error commiting changes, err:", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return err
	}
	return nil
}

// validateEmail validates email re returns error that is safe to share with the user
func (u *User) validateEmail() (statusCode int, err error) {
	res, err := verifier.Verify(u.Email)
	if err != nil {
		fmt.Println("verify email address failed, error is: ", err)
		return http.StatusInternalServerError, err
	}
	// check syntax, needs @ and . for starters
	if !res.Syntax.Valid {
		err = errors.New("email address syntax is invalid")
		fmt.Println(err)
		return http.StatusBadRequest, err
	}
	// check if disposable
	if res.Disposable {
		err = errors.New("sorry, we do not accept disposable email addresses")
		return http.StatusBadRequest, err
	}
	// check if there is domain Suggestion
	if res.Suggestion != "" {
		err = errors.New("email address is not reachable, looking for " + res.Suggestion + " instead?")
		return http.StatusBadRequest, err
	}
	// possible return string values: yes, no, unkown
	if res.Reachable == "no" {
		err = errors.New("email address is not reachable")
		return http.StatusBadRequest, err
	}
	// check MX records so we know DNS setup properly to recieve emails
	if !res.HasMxRecords {
		err = errors.New("domain entered not properly setup to recieve emails, MX record not found")
		return http.StatusBadRequest, err
	}
	return http.StatusOK, nil
}

func (u *User) MakeActive() error {
	stmt, err := db.Prepare("UPDATE users SET is_active = true WHERE id = ?")
	if err != nil {
		fmt.Println("error preparing statement to update is_active")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.ID)
	if err != nil {
		fmt.Println("error executing statement to update is_active")
		return err
	}
	return nil
}

func (u *User) NewEmailVerPswd() (emailVerPswd string, emailVerHash string, err error) {
	// create random code for email
	// Go rune data type represent Unicode characters
	var alphaNumRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	emailVerRandRune := make([]rune, 64)
	// creat a random slice of runes (characters) to create our emailVerPassword (random string of characters)
	for i := 0; i < 64; i++ {
		emailVerRandRune[i] = alphaNumRunes[rand.Intn(len(alphaNumRunes)-1)]
	}
	emailVerPassword := string(emailVerRandRune)
	var emailVerPWhash []byte
	// func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	emailVerPWhash, err = bcrypt.GenerateFromPassword([]byte(emailVerPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt err:", err)
		return emailVerPassword, emailVerHash, err
	}
	emailVerHash = string(emailVerPWhash)
	u.verHash = string(emailVerPWhash)
	return emailVerPassword, emailVerHash, err
}

func (u *User) UpdatePswdHash() error {
	stmt, err := db.Prepare("UPDATE users SET pswd_hash = ? WHERE id = ?")
	if err != nil {
		fmt.Println("error preparing statement to update is_active")
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.pswdHash, u.ID)
	if err != nil {
		fmt.Println("error executing statement to update is_active")
		return err
	}
	return nil
}

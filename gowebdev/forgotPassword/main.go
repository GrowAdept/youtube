package main

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"math/rand"
	"net"
	"net/http"
	"net/smtp"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode"

	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
)

var tpl *template.Template
var db *sql.DB

type TempData struct {
	Username   string
	Email      string
	AuthInfo   string
	ErrMessage string
	Message    string
}

// func NewCookieStore(keyPairs ...[]byte) *CookieStore
var store = sessions.NewCookieStore([]byte(os.Getenv("CookieStorePassword")))

func main() {
	tpl, _ = template.ParseGlob("templates/*.html")
	var err error
	// default internal output of MySQL DATE and DATETIME is []byte, use DSN parameter parseTime=true to use time.Time
	db, err = sql.Open("mysql", "root:"+os.Getenv("MYSQL_PASSWORD")+"@tcp(localhost:3306)/testdb?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/registerval", registerValHandler)
	http.HandleFunc("/registeremailver", registerEmailVerHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/loginver", loginVerHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/forgotpw", forgetpPWhandler)            // renders form for user to enter email address
	http.HandleFunc("/forgotpwval", forgotPWvalHandler)       // validate email address and send email
	http.HandleFunc("/forgotpwchange", forgotpwChangeHandler) // renders change pw form and places authInfo in form action
	http.HandleFunc("/forgotpwemailver", forgotPWverHandler)  // verifies user and resets passwords
	http.HandleFunc("/about", Auth(aboutHandler))
	http.HandleFunc("/index", Auth(indexHandler))
	// if you are not using gorilla/mux, you need to wrap your handler with context.ClearHandler
	http.ListenAndServe("localhost:8080", context.ClearHandler(http.DefaultServeMux))
}

// Auth() adds authentication code to handler before returning handler
// (adds code to check if user is logged in)
// func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
func Auth(HandlerFunc http.HandlerFunc) http.HandlerFunc {
	fmt.Println("*****Auth middleware running*****")
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")
		_, ok := session.Values["userID"]
		if !ok {
			http.Redirect(w, r, "/login", 302)
			return
		}
		// ServeHTTP calls f(w, r)
		// func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
		HandlerFunc.ServeHTTP(w, r)
	}
}

// check session for logged in done with middleware Auth()
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****indexHandler running*****")
	var td TempData
	td.Message = "Logged in"
	tpl.ExecuteTemplate(w, "index.html", td)
}

// check session for logged in done with middleware Auth()
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****aboutHandler running*****")
	var td TempData
	td.Message = "Logged in"
	fmt.Fprint(w, "About Page")
}

// loginHandler serves form for users to login with
func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****loginHandler running*****")
	var td TempData
	tpl.ExecuteTemplate(w, "login.html", td)
}

// loginVerHandler authenticates user login
func loginVerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****loginVerHandler running*****")
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	fmt.Println("username:", username, "password:", password)
	var td TempData
	td.ErrMessage = "check username and password are correct (email must be verified to login)"
	// retrieve password from db to compare (hash) with user supplied password's hash
	var userID, email, dbHash string
	var isActive int
	stmt := "SELECT id, email, hash, is_active FROM users WHERE username = ?"
	row := db.QueryRow(stmt, username)
	err := row.Scan(&userID, &email, &dbHash, &isActive)
	fmt.Println("hash from db:", dbHash)
	if err != nil {
		fmt.Println("error selecting Hash in db by Username, err:", err)
		tpl.ExecuteTemplate(w, "login.html", td)
		return
	}
	// func CompareHashAndPassword(hashedPassword, password []byte) error
	// CompareHashAndPassword() returns err with a value of nil for a match
	err = bcrypt.CompareHashAndPassword([]byte(dbHash), []byte(password))
	if err != nil {
		fmt.Println("couldn't login, incorrect password")
		tpl.ExecuteTemplate(w, "login.html", td)
	}
	// check if user is active, otherwise ask user to verify email
	if isActive == 0 {
		fmt.Println("user is not active yet")
		td.ErrMessage = "Email verification is required before logging in, please check your email"
		tpl.ExecuteTemplate(w, "login.html", td)
		return
	}
	// Get always returns a session, even if empty
	// returns error if exists and could not be decoded
	// Get(r *http.Request, name string) (*Session, error)
	session, _ := store.Get(r, "session")
	// session struct has field make(map[interface{}]interface{})
	session.Values["userID"] = userID
	session.Values["username"] = username
	// save before writing to response/return from handler
	session.Save(r, w)
	td.Message = "Logged in"
	tpl.ExecuteTemplate(w, "index.html", td)
	return
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****logoutHandler running*****")
	session, err := store.Get(r, "session")
	if err != nil {
		fmt.Println("error logging out, err:", err)
		return
	}
	// The delete built-in function deletes the element with the specified key (m[key]) from the map.
	// If m is nil or there is no such element, delete is a no-op.
	delete(session.Values, "userID")
	delete(session.Values, "username")
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		fmt.Println("error logging out, err:", err)
		return
	}
	var td TempData
	td.Message = "Logged out"
	tpl.ExecuteTemplate(w, "login.html", td)
}

// registerHandler serves form for registring new users
func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****registerHandler running*****")
	tpl.ExecuteTemplate(w, "register.html", nil)
}

// registerValHandler creates new user in database
func registerValHandler(w http.ResponseWriter, r *http.Request) {
	/*
		1. check username criteria
		2. check userPassword criteria
		3. check if username already exists in database
		4. create bcrypt hash from userPassword
		5. insert username and userPassword hash in database
		6. check email criteria and check server responds
		7. create emailVerPassword and send in email
		8. save emailVerPWhash in databaser
		9. serve verifyemail.html page
	*/
	fmt.Println("*****registerValHandler running*****")
	r.ParseForm()
	// check username criteria
	username := r.FormValue("username")
	err := checkUsernameCriteria(username)
	if err != nil {
		tpl.ExecuteTemplate(w, "register.html", err.Error())
		return
	}
	// check userPassword criteria
	userPassword := r.FormValue("password")
	err = checkPasswordCriteria(userPassword)
	if err != nil {
		tpl.ExecuteTemplate(w, "register.html", err.Error())
		return
	}
	// check email is valid
	email := r.FormValue("email")
	err = checkEmailValid(email)
	if err != nil {
		tpl.ExecuteTemplate(w, "register.html", err.Error())
		return
	}
	fmt.Println("email:", email, "is valid")
	//check if email domain exists
	err = checkEmailDomain(email)
	if err != nil {
		tpl.ExecuteTemplate(w, "register.html", err.Error())
		return
	}
	var td TempData
	td.ErrMessage = "Sorry, there was a problem registering account, please try again"
	// begin transaction, every query runs successfully or else no changes are made to the database
	// func (db *DB) Begin() (*Tx, error)
	var tx *sql.Tx
	tx, err = db.Begin()
	if err != nil {
		fmt.Println("failed to begin transaction, err:", err)
		tpl.ExecuteTemplate(w, "register.html", td)
		return
	}
	// rollback will be ignored if the tx has been committed later in the function
	defer tx.Rollback()
	// check if username already exists for availability
	stmt := "SELECT id FROM users WHERE username = ?"
	row := tx.QueryRow(stmt, username)
	var uID string
	err = row.Scan(&uID)
	// we only want sql.ErrNoRows, anything else means it already exists or we encountered an error
	if err != sql.ErrNoRows {
		fmt.Println("username exists, err:", err)
		// this message leaves username open to enumeration but usernames are generally accepted at public facing
		td.ErrMessage = "Sorry that username is unavailable"
		tpl.ExecuteTemplate(w, "register.html", td)
		// func (tx *Tx) Rollback() error
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return
	}
	// create userPasswordHashh from userPassword
	var userPasswordHash []byte
	// func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	userPasswordHash, err = bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt err:", err)
		tpl.ExecuteTemplate(w, "register.html", td.ErrMessage)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return
	}
	fmt.Println("userPasswordHash:", userPasswordHash)
	userPasswordHashStr := string(userPasswordHash)
	fmt.Println("userPasswordHashStr:", userPasswordHashStr)
	// insert data into users table
	// func (db *DB) Prepare(query string) (*Stmt, error)
	var insertUserStmt *sql.Stmt
	insertUserStmt, err = tx.Prepare("INSERT INTO users (username, email, hash, created_at, is_active) VALUES (?, ?, ?, ?, ?);")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		tpl.ExecuteTemplate(w, "register.html", td.ErrMessage)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return
	}
	defer insertUserStmt.Close()
	currentTime := time.Now()
	fmt.Println("currentTime:", currentTime)
	var result sql.Result
	//  func (s *Stmt) Exec(args ...interface{}) (Result, error)
	// result, err = insertUserStmt.Exec(username, email, userPasswordHash, currentTime, 0)
	result, err = insertUserStmt.Exec(username, email, userPasswordHashStr, currentTime, 0)
	fmt.Println("err:", err)
	rowsAff, _ := result.RowsAffected()
	fmt.Println("rowsAff:", rowsAff)
	fmt.Println("err:", err)
	// check for successfull insert
	if err != nil || rowsAff != 1 {
		fmt.Println("error inserting new user, err:", err)
		tpl.ExecuteTemplate(w, "register.html", td.ErrMessage)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return
	}
	// create random code for email
	rand.Seed(time.Now().UnixNano())
	// Go rune data type represent Unicode characters
	var alphaNumRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	emailVerRandRune := make([]rune, 64)
	// creat a random slice of runes (characters) to create our emailVerPassword (random string of characters)
	for i := 0; i < 64; i++ {
		emailVerRandRune[i] = alphaNumRunes[rand.Intn(len(alphaNumRunes)-1)]
	}
	fmt.Println("emailVerRandRune:", emailVerRandRune)
	emailVerPassword := string(emailVerRandRune)
	fmt.Println("emailVerPassword:", emailVerPassword)
	var emailVerPWhash []byte
	// func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	emailVerPWhash, err = bcrypt.GenerateFromPassword([]byte(emailVerPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt err:", err)
		tpl.ExecuteTemplate(w, "register.html", td.ErrMessage)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return
	}
	fmt.Println("emailVerPWhash:", emailVerPWhash)
	emailVerPWhashStr := string(emailVerPWhash)
	fmt.Println("emailVerPWhashStr:", emailVerPWhashStr)
	var insertEmailVerStmt *sql.Stmt
	insertEmailVerStmt, err = tx.Prepare("INSERT INTO email_ver_hash (username, email, ver_hash, timeout) VALUES (?, ?, ?, ?);")
	fmt.Println("err:", err)
	if err != nil {
		fmt.Println("error preparing statement:", err)
		tpl.ExecuteTemplate(w, "register.html", td.ErrMessage)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return
	}
	defer insertEmailVerStmt.Close()
	// create timeout limit to register email, 24 hours
	timeout := time.Now().Local().AddDate(0, 0, 1)
	//  func (s *Stmt) Exec(args ...interface{}) (Result, error)
	// result, err = insertEmailVerStmt.Exec(username, email, emailVerPWhash, timeout)
	result, err = insertEmailVerStmt.Exec(username, email, emailVerPWhashStr, timeout)
	rowsAff, _ = result.RowsAffected()
	fmt.Println("rowsAff:", rowsAff)
	fmt.Println("err:", err)
	// check for successfull inserting into email_ver
	if err != nil || rowsAff != 1 {
		fmt.Println("error inserting into email_ver, err:", err)
		tpl.ExecuteTemplate(w, "register.html", td.ErrMessage)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return
	}
	// send emailVerPassword in hyperlink inside email body
	err = sendVerPWemail(emailVerPassword, username, email)
	if err != nil {
		fmt.Println("error emailing hash string, err:", err)
		tpl.ExecuteTemplate(w, "register.html", td.ErrMessage)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return
	}
	// func (tx *Tx) Commit() error
	if commitErr := tx.Commit(); commitErr != nil {
		fmt.Println("error commiting changes, err:", err)
		tpl.ExecuteTemplate(w, "register.html", td.ErrMessage)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return
	}
	fmt.Println("successful insert queries and sent email, committing changes")
	tpl.ExecuteTemplate(w, "verifyemail.html", nil)
}

// registerEmailVerHandler serves form for registring new users
func registerEmailVerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****registerEmailVerHandler running*****")
	username := r.FormValue("u")
	emailVerPassword := r.FormValue("evpw")
	fmt.Println("username:", username)
	fmt.Println("emailVerPassword:", emailVerPassword)
	var td TempData
	td.ErrMessage = "Sorry, there was an issue verifying email, please try again"
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("failed to begin transaction, err:", err)
		tpl.ExecuteTemplate(w, "verifyemail.html", td)
		return
	}
	// rollback will be ignored if the tx has been committed later in the function
	defer tx.Rollback()
	stmt := "SELECT email, ver_hash, timeout FROM email_ver_hash WHERE username = ?"
	row := tx.QueryRow(stmt, username)
	var email string
	var dbEmailVerHash string
	var timeout time.Time
	err = row.Scan(&email, &dbEmailVerHash, &timeout)
	// we only want sql.ErrNoRows, anything else means it already exists or we encountered an error
	if err != nil {
		fmt.Println("error selecting from user_ver_hash, err:", err)
		tpl.ExecuteTemplate(w, "verifyemail.html", td)
		// func (tx *Tx) Rollback() error
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return
	}
	fmt.Println("email:", email)
	fmt.Println("dbEmailVerHash:", dbEmailVerHash)
	fmt.Println("timeout:", timeout)
	// check timeout is within time limit
	currentTime := time.Now()
	if currentTime.After(timeout) {
		fmt.Println("account didn't verify within 24 hours")
		tpl.ExecuteTemplate(w, "verifyemail.html", td)
		// func (tx *Tx) Rollback() error
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbEmailVerHash), []byte(emailVerPassword))
	if err != nil {
		fmt.Println("dbEmailVerHash and hash of emailVerPassword are not the same")
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		tpl.ExecuteTemplate(w, "verifyemail.html", td)
		return
	}
	// ver_hash is a match, setting account is_active to 1 (true)"
	fmt.Println("dbEmailVerHash and hash of emailVerPassword are the same")
	stmt = "UPDATE users SET is_active = 1 WHERE email = ?"
	updateIsActiveStmt, err := tx.Prepare(stmt)
	if err != nil {
		fmt.Println("error preparing updateIsActiveStmt err:", err)
		tpl.ExecuteTemplate(w, "verifyemail.html", td)
		return
	}
	defer updateIsActiveStmt.Close()
	var result sql.Result
	result, err = updateIsActiveStmt.Exec(email)
	rowsAff, _ := result.RowsAffected()
	fmt.Println("rowsAff:", rowsAff)
	// check for successfull insert
	if err != nil || rowsAff != 1 {
		fmt.Println("error inserting new user, err:", err)
		tpl.ExecuteTemplate(w, "verifyemail.html", td.ErrMessage)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return
	}
	if commitErr := tx.Commit(); commitErr != nil {
		fmt.Println("there was an error commiting changes, commitErr:", commitErr)
		tpl.ExecuteTemplate(w, "verifyemail.hmtl", td)
		return
	}
	td.Message = "email verified, go ahead and login"
	td.ErrMessage = ""
	tpl.ExecuteTemplate(w, "login.html", td)
}

// route: /forgotpw
// renders form for recovering account via email
func forgetpPWhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****forgetpPWhandler running*****")
	tpl.ExecuteTemplate(w, "forgotpassword.html", nil)
}

// route:  /forgotpwval
// checks if email exists for this account and send email if exists (do not want to send unsolicited emails)
func forgotPWvalHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****forgotPWvalHandler running*****")
	r.ParseForm()
	email := r.FormValue("email")
	fmt.Println("email:", email)
	var td TempData
	td.ErrMessage = "Sorry, there was an issue recovering account, please try again"
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("failed to begin transaction, err:", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		tpl.ExecuteTemplate(w, "forgotpassword.html", td.ErrMessage)
		return
	}
	// rollback will be ignored if the tx has been committed later in the function
	defer tx.Rollback()
	var username string
	// check email address is in db (do not want to send unsolicited emails)
	row := db.QueryRow("SELECT email, username FROM users WHERE email = ?", email)
	err = row.Scan(&email, &username)
	if err != nil {
		fmt.Println("email not found in db")
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		// notice vague message returned regarding "email address not found", do this if you wish to avoid enumeration
		// if enumeration is not a concern, obvioulsy send a "email address not found" message to user
		tpl.ExecuteTemplate(w, "forgotpwemail.html", nil)
		return
	}
	// create timeout limit
	now := time.Now()
	// add 45 minutes
	timeout := now.Add(time.Minute * 45)
	// Seed uses the provided seed value to initialize the default Source to a deterministic state
	rand.Seed(time.Now().UnixNano())
	// since we are creating this code we aren't forced to make it as easy to guess as a user's password
	var alphaNumRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	emailVerRandRune := make([]rune, 64)
	for i := 0; i < 64; i++ {
		// Intn() returns, as an int, a non-negative pseudo-random number in [0,n).
		emailVerRandRune[i] = alphaNumRunes[rand.Intn(len(alphaNumRunes)-1)]
	}
	fmt.Println("change password emailVerRandRune:", emailVerRandRune)
	emailVerPassword := string(emailVerRandRune)
	fmt.Println("emailVerPassword:", emailVerPassword)
	fmt.Println("emailVerPassword len:", len(emailVerPassword))
	var emailVerPWhash []byte
	// generate emailVerPassword hash for db
	// func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	emailVerPWhash, err = bcrypt.GenerateFromPassword([]byte(emailVerPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt err:", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		tpl.ExecuteTemplate(w, "forgotpassword.html", td.ErrMessage)
		return
	}
	var updateEmailVerStmt *sql.Stmt
	updateEmailVerStmt, err = tx.Prepare("UPDATE email_ver_hash SET ver_hash = ?, timeout = ? WHERE email = ?;")
	if err != nil {
		fmt.Println("error preparing statement:", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		tpl.ExecuteTemplate(w, "forgotpassword.html", td.ErrMessage)
		return
	}
	defer updateEmailVerStmt.Close()
	emailVerPWhashStr := string(emailVerPWhash)
	var result sql.Result
	//  func (s *Stmt) Exec(args ...interface{}) (Result, error)
	result, err = updateEmailVerStmt.Exec(emailVerPWhashStr, timeout, email)
	fmt.Println("err:", err)
	rowsAff, _ := result.RowsAffected()
	fmt.Println("rowsAff:", rowsAff)
	fmt.Println("err:", err)
	// check for successfull insert
	if err != nil || rowsAff != 1 {
		fmt.Println("error inserting new user, err:", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		tpl.ExecuteTemplate(w, "forgotpassword.html", td.ErrMessage)
		return
	}
	// send email with hyperlink
	// sender data
	from := os.Getenv("FromEmailAddr") //ex: "John.Doe@gmail.com"
	password := os.Getenv("SMTPpwd")   // ex: "ieiemcjdkejspqz"
	// receiver address privided through toEmail argument
	to := []string{email}
	// smtp - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	// message
	subject := "Subject: MySite Account Recovery\n"
	// localhost:8080 will be removed by many email service but works with online sites
	// https must be used since we are sending personal data through url parameters
	body := "<body><a rel=\"nofollow noopener noreferrer\" target=\"_blank\" href=\"https://www.mysite.com/forgotpwchange?u=" + username + "&evpw=" + emailVerPassword + "\">Change Password</a></body>"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + body)
	// athentication data
	// func PlainAuth(identity, username, password, host string) Auth
	auth := smtp.PlainAuth("", from, password, host)
	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	fmt.Println("message:", string(message))
	err = smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		fmt.Println("error sending reset password email, err:", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		tpl.ExecuteTemplate(w, "forgotpassword.html", td.ErrMessage)
		return
	}
	if commitErr := tx.Commit(); commitErr != nil {
		fmt.Println("there was an error commiting changes, commitErr:", commitErr)
		tpl.ExecuteTemplate(w, "forgotpassword.html", td.ErrMessage)
	}
	tpl.ExecuteTemplate(w, "forgotpwemail.html", nil)
}

// route: /forgotpwchange
// render form to change password and passes auth info into the path in the form's action
func forgotpwChangeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****forgotpwChangeHandler running*****")
	username := r.FormValue("u")
	emailVerPassword := r.FormValue("evpw")
	fmt.Println("username:", username)
	fmt.Println("emailVerPassword:", emailVerPassword)
	var td TempData
	td.AuthInfo = "?u=" + username + "&evpw=" + emailVerPassword
	tpl.ExecuteTemplate(w, "forgotpwchange.html", td)
}

// route: /forgotpwemailver
// verify emailVerPassword and update db
func forgotPWverHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****forgotPWverHandler running*****")
	username := r.FormValue("u")
	emailVerPassword := r.FormValue("evpw")
	userPassword := r.FormValue("password")
	confirmPassword := r.FormValue("confirmpassword")
	fmt.Println("username:", username)
	fmt.Println("emailVerPassword:", emailVerPassword)
	fmt.Println("userPassword:", userPassword)
	fmt.Println("confirmPassword:", confirmPassword)
	var td TempData
	td.ErrMessage = "Sorry, there was an issue recovering account, please try again"
	td.AuthInfo = "?u=" + username + "&evpw=" + emailVerPassword
	// check if userPassword and confirmpassword are the same
	if userPassword != confirmPassword {
		fmt.Println("passwords do no match")
		td.ErrMessage = "passwords must match"
		tpl.ExecuteTemplate(w, "emailrecoverypw.html", td)
	}
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("failed to begin transaction, err:", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		tpl.ExecuteTemplate(w, "forgotpassword.html", td.ErrMessage)
		return
	}
	// rollback will be ignored if the tx has been committed later in the function
	defer tx.Rollback()
	// retrieving ver_hash and timeout from email_ver_hash table
	var dbEmailVerHash string
	var timeout time.Time
	row := db.QueryRow("SELECT ver_hash, timeout FROM email_ver_hash WHERE username = ?", username)
	err = row.Scan(&dbEmailVerHash, &timeout)
	if err != nil {
		fmt.Println("ver_hash not found in db")
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		tpl.ExecuteTemplate(w, "forgotpassword.html", td.ErrMessage)
		return
	}
	// check if within timelimit
	currentTime := time.Now()
	// func (t Time) After(u Time) bool, After reports whether the time instant t is after u.
	if currentTime.After(timeout) {
		fmt.Println("users:", username, "didn't verify account within 24 hours")
		tpl.ExecuteTemplate(w, "forgotpassword.html", td.ErrMessage)
		// func (tx *Tx) Rollback() error
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return
	}
	fmt.Println("dbEmailVerHash:", dbEmailVerHash)
	// check if db ver_hash is the same as the hash of emailVerPassword from email
	err = bcrypt.CompareHashAndPassword([]byte(dbEmailVerHash), []byte(emailVerPassword))
	if err != nil {
		fmt.Println("dbEmailVerHash and hash of emailVerPassword are not the same")
		// func (tx *Tx) Rollback() error
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		tpl.ExecuteTemplate(w, "forgotpassword.html", td.ErrMessage)
		return
	}
	fmt.Println("dbEmailVerHash and hash of emailVerPassword are the same")
	// check userPassword criteria
	err = checkPasswordCriteria(userPassword)
	if err != nil {
		td.AuthInfo = "?u=" + username + "&evpw=" + emailVerPassword
		// saving password criteria error to inform user
		td.ErrMessage = err.Error()
		tpl.ExecuteTemplate(w, "forgotpwchange.html", td)
		return
	}
	// generate hash for new userPassword
	var hash []byte
	// generate emailVerPassword hash for db
	hash, err = bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt err:", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		tpl.ExecuteTemplate(w, "register.html", td.ErrMessage)
		return
	}
	// update db with new userPasswordHash
	stmt := "UPDATE users SET hash = ? WHERE username = ?"
	updateHashStmt, err := tx.Prepare(stmt)
	if err != nil {
		fmt.Println("error preparing updateHashStmt err:", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		tpl.ExecuteTemplate(w, "forgotpassword.html", td.ErrMessage)
		return
	}
	defer updateHashStmt.Close()
	var result sql.Result
	result, err = updateHashStmt.Exec(hash, username)
	rowsAff, _ := result.RowsAffected()
	fmt.Println("rowsAff:", rowsAff)
	// check for successfull insert
	if err != nil || rowsAff != 1 {
		fmt.Println("error inserting new user, err:", err)
		tpl.ExecuteTemplate(w, "verifyemail.html", td.ErrMessage)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return
	}
	if commitErr := tx.Commit(); commitErr != nil {
		fmt.Println("there was an error commiting changes, commitErr:", commitErr)
		tpl.ExecuteTemplate(w, "forgotpassword.html", td.ErrMessage)
		return
	}
	fmt.Println("forgotten password has been reset")
	td.Message = "Password Successfully Updated"
	tpl.ExecuteTemplate(w, "index.html", td)
}

func sendVerPWemail(randStr string, username string, toEmail string) error {
	// sender data
	from := os.Getenv("FromEmailAddr") //ex: "John.Doe@gmail.com"
	password := os.Getenv("SMTPpwd")   // ex: "ieiemcjdkejspqz"
	// receiver address privided through toEmail argument
	to := []string{toEmail}
	// smtp - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	// message
	subject := "Subject: Email Verification\n"
	// localhost:8080 will be removed by many email service but works with online sites
	// https must be used since we are sending personal data through url parameters
	body := "<body><a rel=\"nofollow noopener noreferrer\" target=\"_blank\" href=\"https://www.mysite.com/registeremailver?u=" + username + "&evpw=" + randStr + "\">mysite</a></body>"
	fmt.Println("body:", body)
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message := []byte(subject + mime + body)
	// athentication data
	// func PlainAuth(identity, username, password, host string) Auth
	auth := smtp.PlainAuth("", from, password, host)
	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	fmt.Println("message:", string(message))
	err := smtp.SendMail(address, auth, from, to, message)
	return err
}

func checkUsernameCriteria(username string) error {
	// check username for only alphaNumeric characters
	var nameAlphaNumeric = true
	for _, char := range username {
		// func IsLetter(r rune) bool, func IsNumber(r rune) bool
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			nameAlphaNumeric = false
		}
	}
	if nameAlphaNumeric != true {
		// func New(text string) error
		return errors.New("Username must only contain letters and numbers")
	}
	// check username length
	var nameLength bool
	if 5 <= len(username) && len(username) <= 50 {
		nameLength = true
	}
	if nameLength != true {
		return errors.New("Username must be longer than 4 characters and less than 51")
	}
	return nil
}

func checkPasswordCriteria(password string) error {
	var err error
	// variables that must pass for password creation criteria
	var pswdLowercase, pswdUppercase, pswdNumber, pswdSpecial, pswdLength, pswdNoSpaces bool
	pswdNoSpaces = true
	for _, char := range password {
		switch {
		// func IsLower(r rune) bool
		case unicode.IsLower(char):
			pswdLowercase = true
		// func IsUpper(r rune) bool
		case unicode.IsUpper(char):
			pswdUppercase = true
			err = errors.New("Pa")
		// func IsNumber(r rune) bool
		case unicode.IsNumber(char):
			pswdNumber = true
		// func IsPunct(r rune) bool, func IsSymbol(r rune) bool
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			pswdSpecial = true
		// func IsSpace(r rune) bool, type rune = int32
		case unicode.IsSpace(int32(char)):
			pswdNoSpaces = false
		}
	}
	// check password length
	if 11 < len(password) && len(password) < 60 {
		pswdLength = true
	}
	// create error for any criteria not passed
	if !pswdLowercase || !pswdUppercase || !pswdNumber || !pswdSpecial || !pswdLength || !pswdNoSpaces {
		switch false {
		case pswdLowercase:
			err = errors.New("Password must contain atleast one lower case letter")
		case pswdUppercase:
			err = errors.New("Password must contain atleast one uppercase letter")
		case pswdNumber:
			err = errors.New("Password must contain atleast one number")
		case pswdSpecial:
			err = errors.New("Password must contain atleast one special character")
		case pswdLength:
			err = errors.New("Passward length must atleast 12 characters and less than 60")
		case pswdNoSpaces:
			err = errors.New("Password cannot have any spaces")
		}
		return err
	}
	return nil
}

func checkEmailValid(email string) error {
	// check email syntax is valid
	//func MustCompile(str string) *Regexp
	emailRegex, err := regexp.Compile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if err != nil {
		fmt.Println(err)
		return errors.New("sorry, something went wrong")
	}
	rg := emailRegex.MatchString(email)
	if rg != true {
		return errors.New("Email address is not a valid syntax, please check again")
	}
	// check email length
	if len(email) < 4 {
		return errors.New("Email length is too short")
	}
	if len(email) > 253 {
		return errors.New("Email length is too long")
	}
	return nil
}

func checkEmailDomain(email string) error {
	i := strings.Index(email, "@")
	host := email[i+1:]
	// func LookupMX(name string) ([]*MX, error)
	_, err := net.LookupMX(host)
	if err != nil {
		err = errors.New("Could not find email's domain server, please chack and try again")
		return err
	}
	return nil
}

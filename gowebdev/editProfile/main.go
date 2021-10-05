package main

import (
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"log"
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

type User struct {
	ID          string
	Username    string
	Email       string
	pswdHash    string
	CreatedAt   time.Time
	IsActive    int
	password    string
	verHash     string
	verPassword string
	timeout     time.Time
}

// func NewCookieStore(keyPairs ...[]byte) *CookieStore
var store = sessions.NewCookieStore([]byte(os.Getenv("CookieStorePassword")))

func main() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
	var err error
	// default internal output of MySQL DATE and DATETIME is []byte, use DSN parameter parseTime=true to use time.Time
	// func Open(driverName, dataSourceName string) (*DB, error)
	db, err = sql.Open("mysql", "root:"+os.Getenv("MYSQL_PASSWORD")+"@tcp(localhost:3306)/testdb2?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	// sql.Open just validates its arguments without creating a connection
	// db.Ping will allow us to know on start up if there is a connection issue,
	// otherwise we wouldn't know there was a connection issue until we run our first query
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	http.HandleFunc("/account/new", accNewHandler)
	http.HandleFunc("/account/create", accCreateHandler)
	http.HandleFunc("/account/email/verify", accEmailVerHandler)
	http.HandleFunc("/account/login", accLoginHandler)
	http.HandleFunc("/account/login/verify", accLoginVerHandler)
	http.HandleFunc("/account/logout", accLogoutHandler)
	http.HandleFunc("/account/forgotpw", accForgotPWhandler)              // renders form for user to enter email address
	http.HandleFunc("/account/forgotpw/validate", accForgotPWvalHandler)  // validate email address and send email
	http.HandleFunc("/account/forgotpw/change", accForgotpwChangeHandler) // renders change pw form and places authInfo in form action
	http.HandleFunc("/account/forgotpw/verify", accForgotPWverHandler)    // verifies user and resets passwords
	http.HandleFunc("/account/profile", Auth(accProfileHandler))          // renders profile form
	http.HandleFunc("/account/edit", Auth(accEditHandler))                // renders edit profile form
	http.HandleFunc("/account/edit/verify", Auth(accEditVerHandler))      // verifies password and updates profile
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/index", Auth(indexHandler))
	addr := "localhost:8080"
	fmt.Print("*****Server Listening at ", addr, "*****\n")
	// if you are not using gorilla/mux, you need to wrap your handler with context.ClearHandler
	// ListenAndServe always returns a non-nil error. (this why we don't use err != nil to check)
	log.Fatal(http.ListenAndServe(addr, context.ClearHandler(http.DefaultServeMux)))
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
			http.Redirect(w, r, "/account/login", 302)
			return
		}
		// ServeHTTP calls f(w, r)
		// func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)
		HandlerFunc.ServeHTTP(w, r)
	}
}

// check session for logged in with middleware Auth()
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****indexHandler running*****")
	var td = map[string]string{
		"UserMessage": "Logged in",
	}
	tpl.ExecuteTemplate(w, "index.html", td)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****aboutHandler running*****")
	tpl.ExecuteTemplate(w, "about.html", nil)
}

// accLoginHandler serves form for users to login with
func accLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****accLoginHandler running*****")
	tpl.ExecuteTemplate(w, "login.html", nil)
}

// accLoginVerHandler authenticates user login
func accLoginVerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****accLoginVerHandler running*****")
	r.ParseForm()
	var u User
	u.Username = r.FormValue("username")
	u.password = r.FormValue("password")
	fmt.Println("u.Username:", u.Username, "u.password:", u.password)
	var td = map[string]string{
		"UserMessage": "check username and password are correct (email must be verified to login)",
	}
	// retrieve password from db to compare (hash) with user supplied password's hash
	err := u.SelectByName()
	if err != nil {
		fmt.Println("error selecting by username, err:", err)
		tpl.ExecuteTemplate(w, "login.html", td)
		return
	}
	// verify form password with db password hash
	err = u.verifyPswd()
	if err != nil {
		if err.Error() == "username and password don't match" {
			fmt.Println(err)
			tpl.ExecuteTemplate(w, "login.html", td)
			return
		}
		fmt.Println("user is not active yet")
		td["UserMessage"] = "Email verification is required before logging in, please check your email"
		tpl.ExecuteTemplate(w, "login.html", td)
		return
	}
	// Get always returns a session, even if empty
	// returns error if exists and could not be decoded
	// Get(r *http.Request, name string) (*Session, error)
	session, _ := store.Get(r, "session")
	// session struct has field Values map[interface{}]interface{}
	session.Values["userID"] = u.ID
	// save before writing to response/return from handler
	session.Save(r, w)
	td["UserMessage"] = "Logged in"
	tpl.ExecuteTemplate(w, "index.html", td)
}

func accLogoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****accLogoutHandler running*****")
	session, err := store.Get(r, "session")
	if err != nil {
		fmt.Println("error logging out, err:", err)
		return
	}
	// The delete built-in function deletes the element with the specified key (m[key]) from the map.
	// If m is nil or there is no such element, delete is a no-op.
	delete(session.Values, "userID")
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		fmt.Println("error logging out, err:", err)
		return
	}
	td := map[string]string{
		"UserMessage": "Logged out",
	}
	tpl.ExecuteTemplate(w, "login.html", td)
}

// accNewHandler serves form for registring new users
func accNewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****accNewHandler running*****")
	tpl.ExecuteTemplate(w, "register.html", nil)
}

// accCreateHandler creates new user in database
func accCreateHandler(w http.ResponseWriter, r *http.Request) {
	/*
		1. check username criteria
		2. check userPassword criteria
		3. check if username already exists in database
		4. create bcrypt hash from userPassword
		5. insert user data into database
		6. check email criteria and check server responds
		7. create u.verHash and send in email
		8. save ver_hash in database
		9. serve verifyemail.html page
	*/
	fmt.Println("*****accCreateHandler running*****")
	r.ParseForm()
	var td = make(map[string]string)
	// check username criteria
	var u User
	u.Username = r.FormValue("username")
	err := u.checkUsernameCriteria()
	if err != nil {
		td["UserMessage"] = err.Error()
		tpl.ExecuteTemplate(w, "register.html", td)
		return
	}
	// check userPassword criteria
	u.password = r.FormValue("password")
	err = u.checkPasswordCriteria()
	if err != nil {
		tpl.ExecuteTemplate(w, "register.html", err.Error())
		return
	}
	// check email is valid
	u.Email = r.FormValue("email")
	err = u.checkEmailValid()
	if err != nil {
		tpl.ExecuteTemplate(w, "register.html", err.Error())
		return
	}
	fmt.Println("u.Email:", u.Email, "is valid")
	//check if email domain exists
	err = u.checkEmailDomain()
	if err != nil {
		tpl.ExecuteTemplate(w, "register.html", err.Error())
		return
	}
	td["UserMessage"] = "Sorry, there was a problem registering account, please try again"
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
	err = u.TxExists(tx)
	// we only want sql.ErrNoRows, anything else means it already exists or we encountered an error
	if err != sql.ErrNoRows {
		fmt.Println("username exists, err:", err)
		// this message leaves username open to enumeration but usernames are generally accepted at public facing
		td["UserMessage"] = "Sorry that username is unavailable"
		tpl.ExecuteTemplate(w, "register.html", td)
		// func (tx *Tx) Rollback() error
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return
	}
	// insert data into users table
	err = u.TxCreate(tx)
	fmt.Println("u.verPassword:", u.verPassword, " (inside of accCreateHandler")
	if err != nil {
		fmt.Println("error inserting user into DB, err:", err)
		tpl.ExecuteTemplate(w, "register.html", td)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return
	}
	// send verifciation password in hyperlink inside email body
	subject := "Email Verification"
	// localhost:8080 will be removed by many email service but works with online sites
	// https must be used since we are sending personal data through url parameters
	body := "<a rel=\"nofollow noopener noreferrer\" target=\"_blank\" href=\"https://www.mysite.com/account/email/verify?u=" + u.Username + "&evpw=" + u.verPassword + "\">mysite</a>"
	// func (u *user) SendEmail(subject, body string) error
	err = u.SendEmail(subject, body)
	if err != nil {
		fmt.Println("error emailing hash string, err:", err)
		tpl.ExecuteTemplate(w, "register.html", td)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return
	}
	// func (tx *Tx) Commit() error
	if commitErr := tx.Commit(); commitErr != nil {
		fmt.Println("error commiting changes, err:", err)
		tpl.ExecuteTemplate(w, "register.html", td)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return
	}
	fmt.Println("successful insert queries and sent email, committing changes")
	tpl.ExecuteTemplate(w, "verifyemail.html", nil)
}

// accEmailVerHandler serves form for registring new users
func accEmailVerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****accEmailVerHandler running*****")
	var u User
	u.Username = r.FormValue("u")
	u.verPassword = r.FormValue("evpw")
	fmt.Println("u.username:", u.Username)
	fmt.Println("u.verPassword:", u.verPassword)
	var td = map[string]string{
		"UserMessage": "Sorry, there was an issue verifying email, please try again",
	}
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("failed to begin transaction, err:", err)
		tpl.ExecuteTemplate(w, "verifyemail.html", td)
		return
	}
	// rollback will be ignored if the tx has been committed later in the function
	defer tx.Rollback()
	err = u.SelectByName()
	if err != nil {
		fmt.Println("error selecting from user_ver_hash, err:", err)
		tpl.ExecuteTemplate(w, "verifyemail.html", td)
		// func (tx *Tx) Rollback() error
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return
	}
	// check timeout is within time limit
	currentTime := time.Now()
	if currentTime.After(u.timeout) {
		fmt.Println("account didn't verify within 24 hours")
		tpl.ExecuteTemplate(w, "verifyemail.html", td)
		// func (tx *Tx) Rollback() error
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.verHash), []byte(u.verPassword))
	if err != nil {
		fmt.Println("dbEmailVerHash and hash of u.verPassword are not the same")
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		tpl.ExecuteTemplate(w, "verifyemail.html", td)
		return
	}
	// ver_hash is a match, setting account is_active to 1 (true)"
	fmt.Println("dbEmailVerHash and hash of u.verPassword are the same")
	err = u.TxMakeActive(tx)
	if err != nil {
		fmt.Println("error inserting new user, err:", err)
		tpl.ExecuteTemplate(w, "verifyemail.html", td)
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
	td["UserMessage"] = "email verified, go ahead and login"
	tpl.ExecuteTemplate(w, "login.html", td)
}

// route: /account/forgotpw
// renders form for recovering account via email
func accForgotPWhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****accForgotPWhandler running*****")
	tpl.ExecuteTemplate(w, "forgotpassword.html", nil)
}

// route:  /account/forgotpw/validate
// checks if email exists for this account and send email if exists (do not want to send unsolicited emails)
func accForgotPWvalHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****accForgotPWvalHandler running*****")
	var u User
	r.ParseForm()
	u.Email = r.FormValue("email")
	fmt.Println("u.email:", u.Email)
	var td = map[string]string{
		"UserMessage": "Sorry, there was an issue recovering account, please try again",
	}
	temp := "forgotpassword.html"
	tx, err := db.Begin()
	fmt.Println("error beggining transaction")
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		tpl.ExecuteTemplate(w, temp, td)
		return
	}
	// rollback will be ignored if the tx has been committed later in the function
	defer tx.Rollback()
	// check email address is in db (do not want to send unsolicited emails)
	err = u.SelectByEmail()
	txCheck(err, tx, w, temp, td, "email not found in db")
	// create timeout limit
	now := time.Now()
	// add 45 minutes
	u.timeout = now.Add(time.Minute * 45)
	err = u.NewVerHash()
	txCheck(err, tx, w, temp, td, "becrypt error")
	err = u.TxUpdate(tx)
	txCheck(err, tx, w, temp, td, "error inserting new user")
	// send email with hyperlink
	// sender data
	subject := "MySite Account Recovery"
	fmt.Println("u.Username before sending email:", u.Username)
	body := "<a rel=\"nofollow noopener noreferrer\" target=\"_blank\" href=\"https://www.mysite.com/account/forgotpw/change?u=" + u.Username + "&evpw=" + u.verPassword + "\">Change Password</a>"
	err = u.SendEmail(subject, body)
	txCheck(err, tx, w, temp, td, "error sending reset password email")
	if commitErr := tx.Commit(); commitErr != nil {
		fmt.Println("there was an error commiting changes, commitErr:", commitErr)
		tpl.ExecuteTemplate(w, "forgotpassword.html", td)
		return
	}
	tpl.ExecuteTemplate(w, "forgotpwemail.html", nil)
}

// route: /account/forgotpw/change
// render form to change password and passes auth info into the path in the form's action
func accForgotpwChangeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****accForgotpwChangeHandler running*****")
	var u User
	u.Username = r.FormValue("u")
	u.verPassword = r.FormValue("evpw")
	fmt.Println("u.Username:", u.Username)
	fmt.Println("u.verPassword:", u.verPassword)
	tpl.ExecuteTemplate(w, "forgotpwchange.html", map[string]string{
		"Username":         u.Username,
		"emailVerPassword": u.verPassword,
	})
}

// route: /account/forgotpw/verify
// verify u.verPassword and update db
func accForgotPWverHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****accForgotPWverHandler running*****")
	temp := "forgotpassword.html"
	var u User
	u.Username = r.FormValue("username")
	u.verPassword = r.FormValue("u.verPassword")
	u.password = r.FormValue("password")
	confirmPassword := r.FormValue("confirmpassword")
	fmt.Println("u.username:", u.Username)
	fmt.Println("u.verPassword:", u.verPassword)
	fmt.Println("u.password:", u.password)
	fmt.Println("confirmPassword:", confirmPassword)
	var td map[string]string
	td["UserMessage"] = "Sorry, there was an issue recovering account, please try again"
	td["AuthInfo"] = "?u=" + u.Username + "&evpw=" + u.verPassword
	// check if userPassword and confirmpassword are the same
	if u.password != confirmPassword {
		fmt.Println("passwords do no match")
		td["UserMessage"] = "passwords must match"
		tpl.ExecuteTemplate(w, "emailrecoverypw.html", td)
	}
	tx, err := db.Begin()
	txCheck(err, tx, w, temp, td, "failed to begin transaction")
	// rollback will be ignored if the tx has been committed later in the function
	defer tx.Rollback()
	// retrieving ver_hash and timeout from email_ver_hash table
	err = u.SelectByName()
	txCheck(err, tx, w, temp, td, "error selecting user from db")
	// check if within timelimit
	currentTime := time.Now()
	// func (t Time) After(u Time) bool, After reports whether the time instant t is after u.
	if currentTime.After(u.timeout) {
		fmt.Println("users:", u.Username, "didn't verify account within 24 hours")
		tpl.ExecuteTemplate(w, "forgotpassword.html", td)
		// func (tx *Tx) Rollback() error
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		return
	}
	fmt.Println("u.verHash:", u.verHash)
	// check if db ver_hash is the same as the hash of u.verPassword from email
	err = bcrypt.CompareHashAndPassword([]byte(u.verHash), []byte(u.verPassword))
	if txCheck(err, tx, w, temp, td, "dbEmailVerHash and hash of u.verPassword are not the same") {
		return
	}
	fmt.Println("dbEmailVerHash and hash of u.verPassword are the same")
	// check userPassword criteria
	err = u.checkPasswordCriteria()
	if err != nil {
		td["AuthInfo"] = "?u=" + u.Username + "&evpw=" + u.verPassword
		// saving password criteria error to inform user
		td["UserMessage"] = err.Error()
		tpl.ExecuteTemplate(w, "forgotpwchange.html", td)
		return
	}
	// generate u.verPassword hash for db
	var pswdByteHash []byte
	pswdByteHash, err = bcrypt.GenerateFromPassword([]byte(u.password), bcrypt.DefaultCost)
	u.pswdHash = string(pswdByteHash)
	if txCheck(err, tx, w, "register.html", td, "bcrypt err") {
		return
	}
	// set timeout back since it was already used
	u.timeout = time.Now().Add(time.Hour * -24)
	// update db with new userPasswordHash
	err = u.TxUpdate(tx)
	if txCheck(err, tx, w, temp, td, "error resetting password") {
		return
	}
	if commitErr := tx.Commit(); commitErr != nil {
		fmt.Println("there was an error commiting changes, commitErr:", commitErr)
		tpl.ExecuteTemplate(w, "forgotpassword.html", td)
		return
	}
	fmt.Println("forgotten password has been reset")
	td["UserMessage"] = "Password Successfully Updated"
	tpl.ExecuteTemplate(w, "index.html", td)
}

// accProfileHandler renders Profile page
func accProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****accProfileHandler running*****")
	var User User
	session, _ := store.Get(r, "session")
	User.ID, _ = session.Values["userID"].(string)
	err := User.SelectByID()
	if err != nil {
		td := map[string]string{
			"UserMessage": "There was an issue displaying profile information",
		}
		tpl.ExecuteTemplate(w, "login.html", td)
		return
	}
	fmt.Println("User:", User)
	tpl.ExecuteTemplate(w, "profile.html", User)
}

// accEditHandler render edit profile page
func accEditHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****accEditHandler running*****")
	var User User
	session, _ := store.Get(r, "session")
	User.ID, _ = session.Values["userID"].(string)
	var td = make(map[string]interface{})
	err := User.SelectByID()
	if err != nil {
		td["UserMessage"] = "There was an issue displaying profile information"
		tpl.ExecuteTemplate(w, "login.html", td)
		return
	}
	fmt.Println("User:", User)
	td["User"] = User
	tpl.ExecuteTemplate(w, "editprofile.html", td)
}

// accEditVefHandler render edit profile page
func accEditVerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*****accEditVerHandler running*****")
	// create td (template data) to pass in User data as well as edit profile messages
	var td = make(map[string]interface{})
	td["UserMessage"] = "There was an issue editing profile"
	// save current db user data into oldUser
	var oldUser User
	session, err := store.Get(r, "session")
	if err != nil {
		fmt.Println("store.Get err:", err)
		td["UserMessage"] = "Please login to Edit Profile"
		tpl.ExecuteTemplate(w, "login.html", td)
		return
	}
	oldUser.ID = session.Values["userID"].(string)
	err = oldUser.SelectByID()
	if err != nil {
		fmt.Println("oldUser.SelectByID() err:", err)
		td["UserMessage"] = "Please login to Edit Profile"
		tpl.ExecuteTemplate(w, "login.html", td)
		return
	}
	fmt.Println("oldUser:", oldUser)
	// save updated user data from form into newUser
	var newUser User
	newUser = oldUser
	newUser.Username = r.FormValue("username")
	newUser.Email = r.FormValue("email")
	newUser.password = r.FormValue("password")
	newPassword := r.FormValue("newpassword")
	confirmPassword := r.FormValue("confirmpassword")
	fmt.Println("newUser:", newUser)
	// oldUser assigned to "User" td value for executing template in case a criteria for update is not met
	td["User"] = oldUser
	// check for changes
	if newUser.Username == oldUser.Username && newUser.Email == oldUser.Email && newPassword == "" {
		fmt.Println("nothing changed in editprofile.html form")
		td["UserMessage"] = "No changes entered into form"
		tpl.ExecuteTemplate(w, "editprofile.html", td)
		return
	}
	// verify user's current password
	err = newUser.verifyPswd()
	if err != nil {
		td["UserMessage"] = err.Error()
		tpl.ExecuteTemplate(w, "editprofile.html", td)
		return
	}
	// confirm newPassword and confirmPassword are the same
	if newPassword != confirmPassword {
		fmt.Println("password and confirm password must be the same")
		td["UserMessage"] = "New password and confirm password must match"
		tpl.ExecuteTemplate(w, "editprofile.html", td)
		return
	}
	// check username criteria
	err = newUser.checkUsernameCriteria()
	if err != nil {
		fmt.Println("newUser.checkUsernameCriteria() err:", err)
		td["UserMessage"] = err.Error()
		tpl.ExecuteTemplate(w, "editprofile.html", td)
		return
	}
	// check if new password has been entered
	if newPassword != "" && confirmPassword != "" {
		fmt.Println("checking password criteria")
		fmt.Println("newPassword:", newPassword)
		// check userPassword criteria
		newUser.password = newPassword
		err = newUser.checkPasswordCriteria()
		if err != nil {
			fmt.Println("newUser.checkPasswordCriteria() err:", err)
			td["UserMessage"] = err.Error()
			tpl.ExecuteTemplate(w, "editprofile.html", td)
			return
		}
		fmt.Println("creating new hash for user password")
		// create new password byte hash
		var userPasswordByteHash []byte
		// func GenerateFromPassword(password []byte, cost int) ([]byte, error)
		userPasswordByteHash, err = bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("bcrypt.GenerateFromPassword() err:", err)
			td["UserMessage"] = err.Error()
			tpl.ExecuteTemplate(w, "editprofile.html", td)
			return
		}
		fmt.Println("userPasswordByteHash:", userPasswordByteHash)
		newUser.pswdHash = string(userPasswordByteHash)
	}
	// check email is valid
	err = newUser.checkEmailValid()
	if err != nil {
		fmt.Println("invalid email")
		td["UserMessage"] = err.Error()
		tpl.ExecuteTemplate(w, "editprofile.html", td)
		return
	}
	//check if email domain exists
	err = newUser.checkEmailDomain()
	if err != nil {
		fmt.Println("newUser.checkEmailDomain() err:", err)
		td["UserMessage"] = err.Error()
		tpl.ExecuteTemplate(w, "editprofile.html", td)
		return
	}
	// update User
	err = newUser.Update()
	if err != nil {
		fmt.Println("newUser.Update() err:", err)
		td["UserMessage"] = "There was an issue updating user profile, please tray again"
		tpl.ExecuteTemplate(w, "editprofile.html", td)
		return
	}
	td["User"] = newUser
	fmt.Println("user data successfully updated!")
	td["UserMessage"] = errors.New("Your profile has been updated")
	fmt.Println("td:", td)
	// send email notifying user of changes (incase it wasn't them)
	subject := "Profile Updated"
	body := "<h2>Your profile has been updated.</h2>"
	err = oldUser.SendEmail(subject, body)
	if err != nil {
		fmt.Println("newUser.SendEmail() err:", err)
	}
	tpl.ExecuteTemplate(w, "editprofile.html", td)
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
	body := "<body><a rel=\"nofollow noopener noreferrer\" target=\"_blank\" href=\"https://www.mysite.com/account/email/verify?u=" + username + "&evpw=" + randStr + "\">mysite</a></body>"
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

func (u *User) checkUsernameCriteria() error {
	// check username for only alphaNumeric characters
	var nameAlphaNumeric = true
	for _, char := range u.Username {
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
	if 5 <= len(u.Username) && len(u.Username) <= 50 {
		nameLength = true
	}
	if nameLength != true {
		return errors.New("Username must be longer than 4 characters and less than 51")
	}
	return nil
}

func (u *User) checkPasswordCriteria() error {
	var err error
	// variables that must pass for password creation criteria
	var pswdLowercase, pswdUppercase, pswdNumber, pswdSpecial, pswdLength, pswdNoSpaces bool
	pswdNoSpaces = true
	for _, char := range u.password {
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
	if 11 < len(u.password) && len(u.password) < 60 {
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

func (u *User) checkEmailValid() error {
	// check email syntax is valid
	//func MustCompile(str string) *Regexp
	emailRegex, err := regexp.Compile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if err != nil {
		fmt.Println(err)
		return errors.New("sorry, something went wrong")
	}
	rg := emailRegex.MatchString(u.Email)
	if rg != true {
		return errors.New("Email address is not a valid syntax, please check again")
	}
	// check email length
	if len(u.Email) < 4 {
		return errors.New("Email length is too short")
	}
	if len(u.Email) > 253 {
		return errors.New("Email length is too long")
	}
	return nil
}

func (u *User) checkEmailDomain() error {
	i := strings.Index(u.Email, "@")
	host := u.Email[i+1:]
	// func LookupMX(name string) ([]*MX, error)
	_, err := net.LookupMX(host)
	if err != nil {
		err = errors.New("Could not find email's domain server, please check and try again")
		return err
	}
	return nil
}

func (u *User) SelectByName() error {
	stmt := "SELECT id, email, pswd_hash, created_at, is_active, ver_hash, timeout FROM users WHERE username = ?"
	row := db.QueryRow(stmt, u.Username)
	err := row.Scan(&u.ID, &u.Email, &u.pswdHash, &u.CreatedAt, &u.IsActive, &u.verHash, &u.timeout)
	if err != nil {
		fmt.Println("error selecting user by name, err:", err)
		return err
	}
	return err
}

func (u *User) SelectByID() error {
	stmt := "SELECT id, username, email, pswd_hash, created_at, is_active, ver_hash, timeout FROM users WHERE id = ?"
	row := db.QueryRow(stmt, u.ID)
	err := row.Scan(&u.ID, &u.Username, &u.Email, &u.pswdHash, &u.CreatedAt, &u.IsActive, &u.verHash, &u.timeout)
	if err != nil {
		fmt.Println("error selecting user by id, err:", err)
		return err
	}
	return err
}

func (u *User) SelectByEmail() error {
	stmt := "SELECT id, username, pswd_hash, created_at, is_active, ver_hash, timeout FROM users WHERE email = ?"
	row := db.QueryRow(stmt, u.Email)
	err := row.Scan(&u.ID, &u.Username, &u.pswdHash, &u.CreatedAt, &u.IsActive, &u.verHash, &u.timeout)
	if err != nil {
		fmt.Println("error selecting user by name, err:", err)
		return err
	}
	return err
}

// TxExists check to see if username already exists
func (u *User) TxExists(tx *sql.Tx) error {
	stmt := "SELECT id FROM users WHERE username = ?"
	row := tx.QueryRow(stmt, u.Username)
	err := row.Scan(&u.ID)
	return err
}

func (u *User) TxCreate(tx *sql.Tx) (err error) {
	// func (db *DB) Prepare(query string) (*Stmt, error)
	var insertUserStmt *sql.Stmt
	insertUserStmt, err = tx.Prepare("INSERT INTO users (username, email, pswd_hash, created_at, is_active, ver_hash, timeout) VALUES (?, ?, ?, ?, ?, ?, ?);")
	if err != nil {
		return err
	}
	defer insertUserStmt.Close()
	// generate time for user.createdAt
	u.CreatedAt = time.Now()
	fmt.Println("u.createdAt:", u.CreatedAt)
	// create userPasswordHashh from userPassword
	var userPasswordByteHash []byte
	// func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	userPasswordByteHash, err = bcrypt.GenerateFromPassword([]byte(u.password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("bcrypt err:", err)
		return err
	}
	fmt.Println("userPasswordByteHash:", userPasswordByteHash)
	u.pswdHash = string(userPasswordByteHash)
	fmt.Println("u.pswdHash:", u.pswdHash)
	// create u.verHash
	err = u.NewVerHash()
	fmt.Println("u.verPassword:", u.verPassword, "(inside of u.Create())")
	if err != nil {
		return err
	}
	// create timeout limit to register email, 5 days
	u.timeout = time.Now().Local().AddDate(0, 0, 5)
	var result sql.Result
	// func (s *Stmt) Exec(args ...interface{}) (Result, error)
	// result, err = insertUserStmt.Exec(username, email, userPasswordHash, currentTime, 0)
	result, err = insertUserStmt.Exec(u.Username, u.Email, u.pswdHash, u.CreatedAt, 0, u.verHash, u.timeout)
	fmt.Println("err:", err)
	rowsAff, _ := result.RowsAffected()
	fmt.Println("rowsAff:", rowsAff)
	fmt.Println("err:", err)
	// check for successfull insert
	if err != nil || rowsAff != 1 {
		fmt.Println("error inserting new user, err:", err)
		return errors.New("error inserting new user or affected more than one row")
	}
	return nil
}

func (u *User) NewVerHash() error {
	// Seed uses the provided seed value to initialize the default Source to a deterministic state
	rand.Seed(time.Now().UnixNano())
	// Go rune data type represent Unicode characters
	var alphaNumRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	VerRandRune := make([]rune, 64)
	// creat a random slice of runes (characters) to create our verifcation password for our meail (random string of characters)
	for i := 0; i < 64; i++ {
		VerRandRune[i] = alphaNumRunes[rand.Intn(len(alphaNumRunes)-1)]
	}
	fmt.Println("VerRandRune:", VerRandRune)
	u.verPassword = string(VerRandRune)
	fmt.Println("u.verPassword:", u.verPassword, " (inside of u.NewVerHash() 1st)")
	var VerPWhash []byte
	// func GenerateFromPassword(password []byte, cost int) ([]byte, error)
	VerPWhash, err := bcrypt.GenerateFromPassword([]byte(u.verPassword), bcrypt.DefaultCost)
	fmt.Println("VerPWhash:", VerPWhash)
	u.verHash = string(VerPWhash)
	fmt.Println("string(VerPWhash):", string(VerPWhash))
	fmt.Println("u.verPassword:", u.verPassword, " (inside of u.NewVerHash() 2nd)")
	return err
}

func (u *User) SendEmail(subject, body string) error {
	// sender data
	from := os.Getenv("FromEmailAddr") //ex: "John.Doe@gmail.com"
	password := os.Getenv("SMTPpwd")   // ex: "ieiemcjdkejspqz"
	// receiver address privided through toEmail argument
	to := []string{u.Email}
	// smtp - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject = "Subject: " + subject + "\n"
	body = "<body>" + body + "</body>"
	message := []byte(subject + mime + body)
	// athentication data
	// func PlainAuth(identity, username, password, host string) Auth
	auth := smtp.PlainAuth("", from, password, host)
	// func SendMail(addr string, a Auth, from string, to []string, msg []byte) error
	fmt.Println("message:", string(message))
	err := smtp.SendMail(address, auth, from, to, message)
	fmt.Println("u.Sendmail() err:", err)
	return err
}

func (u *User) TxMakeActive(tx *sql.Tx) error {
	stmt := "UPDATE users SET is_active = 1 WHERE email = ?"
	updateIsActiveStmt, err := tx.Prepare(stmt)
	if err != nil {
		fmt.Println("error preparing updateIsActiveStmt err:", err)
		return errors.New("error preparing query to make user active")
	}
	defer updateIsActiveStmt.Close()
	var result sql.Result
	result, err = updateIsActiveStmt.Exec(u.Email)
	rowsAff, _ := result.RowsAffected()
	fmt.Println("rowsAff:", rowsAff)
	// check for successfull insert
	if rowsAff != 1 {
		fmt.Println("error inserting new user, err:", err)
		return errors.New("rows affected not equal to 1")
	}
	return err
}

func (u *User) TxUpdate(tx *sql.Tx) error {
	var updateUserStmt *sql.Stmt
	var err error
	updateUserStmt, err = tx.Prepare("UPDATE users SET username = ?, email = ?, pswd_hash = ?, is_active = ?, ver_hash = ?, timeout = ? WHERE id = ?;")
	if err != nil {
		fmt.Println("error preparring statement to update user in DB with u.TxUpdate, err:", err)
		return err
	}
	defer updateUserStmt.Close()
	var result sql.Result
	//  func (s *Stmt) Exec(args ...interface{}) (Result, error)
	result, err = updateUserStmt.Exec(u.Username, u.Email, u.pswdHash, u.IsActive, u.verHash, u.timeout, u.ID)
	fmt.Println("err:", err)
	rowsAff, _ := result.RowsAffected()
	fmt.Println("rowsAff:", rowsAff)
	fmt.Println("err:", err)
	// check for successfull update
	if rowsAff != 1 {
		fmt.Println("rows affected not equal to one:", err)
		return errors.New("number of rows affected not equal to one")
	}
	return err
}

func (u *User) Update() error {
	fmt.Println(".Update running")
	var updateUserStmt *sql.Stmt
	var err error
	updateUserStmt, err = db.Prepare("UPDATE users SET username = ?, email = ?, pswd_hash = ?, is_active = ?, ver_hash = ?, timeout = ? WHERE id = ?;")
	if err != nil {
		fmt.Println("error preparring statement to update user in DB with u.Update(), err:", err)
		return err
	}
	defer updateUserStmt.Close()
	var result sql.Result
	//  func (s *Stmt) Exec(args ...interface{}) (Result, error)
	result, err = updateUserStmt.Exec(u.Username, u.Email, u.pswdHash, u.IsActive, u.verHash, u.timeout, u.ID)
	fmt.Println("err:", err)
	fmt.Println("result:", result)
	rowsAff, _ := result.RowsAffected()
	fmt.Println("rowsAff:", rowsAff)
	// check for successfull update
	if err != nil {
		fmt.Println("there was an erorr updating user in u.Update() err:", err)
		return errors.New("number of rows affected not equal to one")
	}
	if rowsAff != 1 {
		fmt.Println("rows affected not equal to one:", err)
		return errors.New("number of rows affected not equal to one")
	}
	return err
}

// txCheck checks CRUD operations for errors and handles with rollback and executing template for user with message on error
// this function is intended to be used with an if statement to return our of function it was called to avoid displaying two templates
func txCheck(err error, tx *sql.Tx, w http.ResponseWriter, temp string, td map[string]string, message string) bool {
	if err != nil {
		fmt.Println(message, ", err:", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			fmt.Println("there was an error rolling back changes, rollbackErr:", rollbackErr)
		}
		tpl.ExecuteTemplate(w, temp, td)
		return true
	}
	return false
}

// verifyPswd checks the user's password hash from the db with the password password given in a form
// user.Select needs to be run to populate fields before running this function
func (u *User) verifyPswd() (err error) {
	// CompareHashAndPassword() returns err with a value of nil for a match
	// func CompareHashAndPassword(hashedPassword, password []byte) error
	err = bcrypt.CompareHashAndPassword([]byte(u.pswdHash), []byte(u.password))
	if err != nil {
		fmt.Println("username and password don't match")
		err = errors.New("username and password don't match")
		return err
	}
	// check if user is active, otherwise ask user to verify email
	if u.IsActive == 0 {
		fmt.Println("user email not verified yet")
		err = errors.New("account not active")
		return
	}
	return
}

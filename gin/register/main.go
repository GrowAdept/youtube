package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

const minEntropyBits = 60

var (
	verifier = emailverifier.NewVerifier()
)

func init() {
	verifier = verifier.EnableDomainSuggest()
	/*
		port 25 is blocked for most at home ISP so EnableSMTPCheck() needs to be blocked
		at home, hosting providers have port 25 open or allow it to be opened
	*/
	if servLoc == "remote" {
		verifier = verifier.EnableSMTPCheck() // enable for remote production server
	}
	dispEmailsDomains := MustDispEmailDom()
	verifier = verifier.AddDisposableDomains(dispEmailsDomains)
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*/*.html")

	router.GET("/", indexHandler)
	router.GET("/login", loginGEThandler)
	router.POST("/login", loginPOSThandler)
	router.GET("/register", registerGEThandler)
	router.POST("/register", registerPOSThandler)
	router.GET("/emailver/:username/:verPass", emailverGEThandler)

	authRouter := router.Group("/user", auth)
	authRouter.GET("/profile", profileHandler)
	authRouter.GET("/logout", logoutHandler)

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}

// index page
func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// loginGEThandler displays form for login
func loginGEThandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

// loginPOSThandler verifies login credentials
func loginPOSThandler(c *gin.Context) {
	var user User
	user.Username = c.PostForm("username")
	password := c.PostForm("password")
	err := user.getUserByUsername()
	if err != nil {
		fmt.Println("error selecting pswd_hash in db by Username, err:", err)
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"message": "check username and password"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.pswdHash), []byte(password))
	if err == nil {
		session, _ := store.Get(c.Request, "session")
		// session struct has field Values map[interface{}]interface{}
		session.Values["user"] = user
		// save before writing to response/return from handler
		session.Save(c.Request, c.Writer)
		c.HTML(http.StatusOK, "loggedin.html", gin.H{"username": user.Username})
		return
	}
	fmt.Println("err:", err)
	c.HTML(http.StatusUnauthorized, "login.html", gin.H{"message": "check username and password"})
}

// logoutHandler logout user by deleting session data
func logoutHandler(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")
	delete(session.Values, "user")
	session.Save(c.Request, c.Writer)
	c.HTML(http.StatusOK, "login.html", gin.H{"message": "Logged out"})
}

// profileHandler displays profile information
func profileHandler(c *gin.Context) {
	session, _ := store.Get(c.Request, "session")
	var user = &User{}
	val := session.Values["user"]
	var ok bool
	if user, ok = val.(*User); !ok {
		fmt.Println("was not of type *User")
		c.HTML(http.StatusForbidden, "login.html", nil)
		return
	}
	c.HTML(http.StatusOK, "profile.html", gin.H{"user": user})
}

func registerGEThandler(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func registerPOSThandler(c *gin.Context) {
	var u User
	u.Username = c.PostForm("username")
	u.Email = c.PostForm("email")
	u.password = c.PostForm("password")
	// validate username
	err := u.validateUsername()
	if err != nil {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"message": err,
			"user":    u,
		})
		return
	}
	// validate password
	err = u.validatePassword()
	if err != nil {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"message": err,
			"user":    u,
		})
		return
	}
	// validate email
	var statusCode int
	statusCode, err = u.validateEmail()
	if err != nil {
		c.HTML(statusCode, "register.html", gin.H{
			"message": err,
			"user":    u,
		})
		return
	}
	// check if username already exists
	exists := u.UsernameExists()
	if exists {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"message": "Username already taken, please try another",
			"user":    u,
		})
		return
	}
	// create user pswd hash and save user data
	err = u.New()
	if err != nil {
		fmt.Println("create.New err:", err)
		err = errors.New("there was an issue creating account, please try again")
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"message": err,
			"user":    u,
		})
		return
	}
	c.HTML(http.StatusOK, "register-succ.html", gin.H{})
}

func emailverGEThandler(c *gin.Context) {
	var u User
	u.Username = c.Param("username")
	linkVerPass := c.Param("verPass")
	err := u.getUserByUsername()
	if err != nil {
		fmt.Println("error selecting ver_hash in db by Username, err:", err)
		c.HTML(http.StatusUnauthorized, "register-succ.html", gin.H{"message": "Please try link in verification email again"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.verHash), []byte(linkVerPass))
	if err == nil {
		// update user.Active to true
		err = u.MakeActive()
		if err != nil {
			c.HTML(http.StatusBadRequest, "acc-activated.html", gin.H{
				"message": "Please try email confirmation link again",
			})
			return
		}
		c.HTML(http.StatusOK, "acc-activated.html", nil)
		return
	}
	c.HTML(http.StatusUnauthorized, "register-succ.html", gin.H{"message": "Please try link in verification email again"})
}

// disposable email list from
// https://github.com/disposable-email-domains/disposable-email-domains/blob/master/disposable_email_blocklist.conf
func MustDispEmailDom() (dispEmailDomains []string) {
	file, err := os.Open("disposable_email_blocklist.txt")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dispEmailDomains = append(dispEmailDomains, scanner.Text())
	}
	return dispEmailDomains
}

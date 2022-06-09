package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"

	"github.com/GrowAdept/youtube/gin/register/models"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var store = sessions.NewCookieStore([]byte("super-secret"))

func init() {
	store.Options.HttpOnly = true // since we are not accessing any cookies w/ JavaScript, set to true
	store.Options.Secure = true   // requires secuire HTTPS connection
	gob.Register(&models.User{})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*/*.html")
	var err error

	router.GET("/", indexHandler)
	router.GET("/login", loginGEThandler)
	router.POST("/login", loginPOSThandler)

	authRouter := router.Group("/user", auth)
	authRouter.GET("/profile", profileHandler)

	err = router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}

// auth middleware
func auth(c *gin.Context) {
	fmt.Println("auth middleware running")
	session, _ := store.Get(c.Request, "session")
	fmt.Println("session:", session)
	_, ok := session.Values["user"]
	if !ok {
		c.HTML(http.StatusForbidden, "login.html", nil)
		c.Abort()
		return
	}
	fmt.Println("middleware done")
	c.Next()
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
	var user models.User
	user.Username = c.PostForm("username")
	password := c.PostForm("password")
	err := user.GetUserByUsername()
	if err != nil {
		fmt.Println("error selecting pswd_hash in db by Username, err:", err)
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{"message": "check username and password"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.pswdHash), []byte(password))
	fmt.Println("err from bycrypt:", err)
	if err == nil {
		session, _ := store.Get(c.Request, "session")
		// session struct has field Values map[interface{}]interface{}
		session.Values["user"] = user
		// save before writing to response/return from handler
		session.Save(c.Request, c.Writer)
		c.HTML(http.StatusOK, "loggedin.html", gin.H{"username": user.Username})
		return
	}
	c.HTML(http.StatusUnauthorized, "login.html", gin.H{"message": "check username and password"})
}

// profileHandler displays profile information
func profileHandler(c *gin.Context) {
	fmt.Println("profile handler running")
	session, _ := store.Get(c.Request, "session")
	var user = &models.User{}
	val := session.Values["user"]
	var ok bool
	if user, ok = val.(*User); !ok {
		fmt.Println("was not of type *User")
		c.HTML(http.StatusForbidden, "login.html", nil)
		return
	}
	c.HTML(http.StatusOK, "profile.html", gin.H{"user": user})
}

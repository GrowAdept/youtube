package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

const minEntropyBits = 60

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/register", registerGetHandler)
	router.POST("/register", registerPostHandler)
	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}

func registerGetHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func registerPostHandler(c *gin.Context) {
	password := c.PostForm("password")

	// entropy is a float64, representing the strength in base 2 (bits)
	entropy := passwordvalidator.GetEntropy(password)
	fmt.Println("entropy:", entropy)

	// if the password has enough entropy, err is nil
	// otherwise, a formatted error message is provided explaining
	// how to increase the strength of the password
	// (safe to show to the client)
	err := passwordvalidator.Validate(password, minEntropyBits)
	fmt.Println("err:", err)
	if err != nil {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"message": err})
		return
	}
	c.HTML(http.StatusOK, "register-succ.html", gin.H{"entropy": entropy})
}

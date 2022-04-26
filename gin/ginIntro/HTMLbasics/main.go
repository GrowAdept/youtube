package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Default returns an Engine instance with the Logger and
	// Recovery middleware already attached.
	// func Default() *Engine
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	// send simple message
	router.GET("/hello", getHello)
	// use our parsed template
	router.GET("/greet", getGreet)
	// return the value of the URL param.
	router.GET("/greet/:name", getGreetName)
	// use gin.H map to pass multiple var into template
	router.GET("/many", getManyData)
	// render template
	router.GET("/form", getForm)
	// retrieve form data
	router.POST("/form", postForm)

	// Run attaches the router to a http.Server and starts listening and serving HTTP requests.
	// func (engine *Engine) Run(addr ...string) (err error)
	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}

// GET /hello
func getHello(c *gin.Context) {
	// String writes the given string into the response body.
	// http.StautsOK is http status code saved as constant in http package
	// func (c *Context) String(code int, format string, values ...interface{})
	c.String(http.StatusOK, "Hello World!")
}


// GET /greet
func getGreet(c *gin.Context) {
	// HTML renders the HTTP template specified by its file name.
	// It also updates the HTTP code and sets the Content-Type as "text/html"
	// func (c *Context) HTML(code int, name string, obj interface{})
	c.HTML(http.StatusOK, "greeting.html", nil)
}

// GET /great/:name
func getGreetName(c *gin.Context) {
	// Param returns the value of the URL param.
	// It is a shortcut for c.Params.ByName(key)
	// func (c *Context) Param(key string) string
	name := c.Param("name")
	c.HTML(http.StatusOK, "customGreeting.html", name)
}

// GET /many
func getManyData(c *gin.Context) {
	foods := []string{"chicken sandwich", "fries", "soda", "cookie"}
	// H is a shortcut for map[string]interface{}
	// type H map[string]interface{}
	c.HTML(http.StatusOK, "manyData.html", gin.H{
		"name":  "Carl",
		"foods": foods,
	})
}

// GET /form
func getForm(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", nil)
}

// POST /form
func postForm(c *gin.Context) {
	// PostForm returns the specified key from a POST urlencoded form or multipart form when it exists,
	// otherwise it returns an empty string `("")`.
	// func (c *Context) PostForm(key string) string
	name := c.PostForm("name")
	food := c.PostForm("food")
	c.HTML(http.StatusOK, "formResult.html", gin.H{
		"name": name,
		"food": food,
	})
}

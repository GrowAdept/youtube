package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Gopher struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var gophers = []Gopher{
	{"1", "Ken", "Thompson"},
	{"2", "Robert", "Griesemer"},
}

func main() {
	router := gin.Default()
	router.GET("/gopher", getGophers)
	router.POST("/gopher", createGopher)
	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}

/*
c.IndentedJSON(http.StatusOK, gopher)
 -automatically sets content type to application/json
 -uses passed in status code
 -serializes given struct as pretty JSON
*/
func getGophers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gophers)
}

/*
c.BindJSON(&newGohper)
 -binds JSON to struct pointer
 -we handle any errors
*/
func createGopher(c *gin.Context) {
	var newGohper Gopher
	err := c.BindJSON(&newGohper)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	gophers = append(gophers, newGohper)
	c.IndentedJSON(http.StatusCreated, gophers)
}

/*
	// BindJSON is a shortcut for c.MustBindWith(obj, binding.JSON).
	// MustBindWith binds the passed struct pointer using the specified binding engine.
	// It will abort the request with HTTP 400 if any error occurs. See the binding package.
	// func (c *Context) BindJSON(obj interface{}) error
*/

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// movie represents data about a film.
type movie struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Price    string `json:"price"`
}

// movies slice for demonstration, really want db for persistance
var movies = []movie{
	{ID: "1", Title: "The Dark Knight", Director: "Christopher Nolan", Price: "5.99"},
	{ID: "2", Title: "Tommy Boy", Director: "Peter Segal", Price: "2.99"},
	{ID: "3", Title: "The Shawshank Redemption", Director: "Frank Darabont", Price: "7.99"},
}

func main() {
	router := gin.Default()

	// API (Application Programming Interface)
	router.GET("/movie", getMovies)              // get all movies
	router.GET("/movie/:id", getMovieByID)       // get one movie by ID
	router.POST("/movie", createMovie)           // create new movie
	router.PATCH("/movie/:id", updateMoviePrice) // update movie price
	router.DELETE("/movie/:id", deleteMovie)     // delete movie

	err := router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}

// using slice in place of using db
func getMovies(c *gin.Context) {
	/*
		IndentedJSON serializes the given struct as pretty JSON (indented + endlines)
		into the response body. It also sets the Content-Type as "application/json".
		WARNING: we recommend to use this only for development purposes since printing
		pretty JSON is more CPU and bandwidth consuming. Use Context.JSON() instead.
		func (c *Context) IndentedJSON(code int, obj interface{})
	*/
	c.IndentedJSON(http.StatusOK, movies)
}

func getMovieByID(c *gin.Context) {
	id := c.Param("id")
	var index int
	// search for particular movie in slice for demo (db in production)
	for k, v := range movies {
		if id == v.ID {
			index = k
		}
	}
	c.IndentedJSON(http.StatusOK, movies[index])
}

func createMovie(c *gin.Context) {
	var newMovie movie
	// BindJSON is a shortcut for c.MustBindWith(obj, binding.JSON).
	// MustBindWith binds the passed struct pointer using the specified binding engine.
	// It will abort the request with HTTP 400 if any error occurs. See the binding package.
	// func (c *Context) BindJSON(obj interface{}) error
	err := c.BindJSON(&newMovie)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	movies = append(movies, newMovie)
	c.IndentedJSON(http.StatusCreated, movies)
}

func updateMoviePrice(c *gin.Context) {
	var index int
	id := c.Param("id")
	for k, v := range movies {
		if id == v.ID {
			index = k
		}
	}
	movies[index].Price = "9.99"
	c.IndentedJSON(http.StatusOK, movies[index])
}

func deleteMovie(c *gin.Context) {
	var index int
	id := c.Param("id")
	for k, v := range movies {
		if id == v.ID {
			index = k
		}
	}
	// delete item from slice
	movies = append(movies[:index], movies[index+1:]...)
	c.IndentedJSON(http.StatusOK, movies)
}

package main

import (
	"fmt"
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
	// gin.Default attaches Logger and Recovery middleware
	// func gin.New() *gin.Engine creates router without middleware
	router := gin.New()
	router.LoadHTMLGlob("templates/*.html")

	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	router.Use(middlewareFunc1, middlewareFunc2, middlewareFunc3())

	router.GET("/movie", getAllMovies) // get all movies
	authRouter := router.Group("/auth", gin.BasicAuth(gin.Accounts{
		"Joe":   "baseball",
		"Kelly": "1234",
	}))

	// path http://localhost:8080/auth/movie
	authRouter.GET("/movie", createMovieForm) // create new movie
	authRouter.POST("/movie", createMovie)    // create new movie

	router.Run(":8080")
}

func middlewareFunc1(c *gin.Context) {
	fmt.Println("middlewareFunc1 running")
	// Next should be used only inside middleware.
	//It executes the pending handlers in the chain inside the calling handler.
	c.Next()
}

func middlewareFunc2(c *gin.Context) {
	fmt.Println("middlewareFunc2 running")
	// Abort prevents pending handlers from being called
	// c.Abort()
	fmt.Println("middlewareFunc2 ending")
	c.Next()
}

func middlewareFunc3() gin.HandlerFunc {
	// run one time logic could inserted here
	return func(c *gin.Context) {
		fmt.Println("middlewareFunc3 running")
		c.Next()
	}
}

func getAllMovies(c *gin.Context) {
	c.HTML(http.StatusOK, "allmovies.html", movies)
}

func createMovieForm(c *gin.Context) {
	c.HTML(http.StatusOK, "createmovieform.html", nil)
}

func createMovie(c *gin.Context) {
	var newMovie movie
	newMovie.ID = c.PostForm("id")
	newMovie.Title = c.PostForm("title")
	newMovie.Director = c.PostForm("director")
	newMovie.Price = c.PostForm("price")
	movies = append(movies, newMovie)
	c.HTML(http.StatusOK, "allmovies.html", movies)
}

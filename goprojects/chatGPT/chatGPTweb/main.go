package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	gpt "github.com/GrowAdept/youtube/goprojects/chatGPT/chatGPTapi"
	"github.com/gin-gonic/gin"
)

var APIkey = os.Getenv("CHATGPT_API_KEY")
var Client = gpt.CreateClient(APIkey, "https://api.openai.com/v1/completions")

func main() {
	answ, err := Client.AskGPTansw("what is a ferret")
	fmt.Println("answ:", answ)
	fmt.Println("err:", err)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/ask", askGEThandler)
	router.POST("/ask", askPOSThandler)

	err = router.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}

// askGetHandler displays field to ask chatGPT questions
func askGEThandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// askPostHandler asks chatGPT question from user and serves response
func askPOSThandler(c *gin.Context) {
	question := c.PostForm("question")
	answ, err := Client.AskGPTansw(question)
	if err != nil {
		fmt.Println("error asking Chat GPT:", err)
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{"message": "sorry, something went wrong"})
		return
	}
	c.HTML(http.StatusAccepted, "index.html", gin.H{"message": answ})
}

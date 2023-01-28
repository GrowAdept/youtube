package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	// C:\Program Files\Go\src\github.com\GrowAdept\youtube\goprojects\chatGPT\chatGPTapi
	// C:\Users\jut55\OneDrive\gocode\src\github.com\GrowAdept\youtube\goprojects\chatGPT\chatGPTapi
	// C:\Users\jut55\OneDrive\gocode\src\youtube\goprojects\chatGPT\chatGPTweb\main.go
	// C:\Users\jut55\OneDrive\gocode
	// C:\Users\jut55\OneDrive\gocode;C:\Users\jut55\OneDrive\gocode\bin
	gpt "github.com/GrowAdept/youtube/goprojects/chatGPT/chatGPTapi"
	"github.com/gin-gonic/gin"
)

var APIkey = os.Getenv("CHATGPT_API_KEY")
var Client = gpt.CreateClient(APIkey, "https://api.openai.com/v1/completions")

type PromptParams2 struct {
	Model            string         `json:"model"` // required
	Prompt           string         `json:"prompt"`
	Suffix           string         `json:"suffix,omitempty"`
	MaxTokens        int            `json:"max_tokens,omitempty"`
	Temperature      float64        `json:"temperature"`
	TopP             float64        `json:"top_p"`
	N                int            `json:"n"`
	Stream           bool           `json:"stream"`
	LogProbs         int            `json:"logprobs"`
	Echo             bool           `json:"echo"`
	Stop             string         `json:"stop"`
	PressencePenalty float64        `json:"presence_penalty"`
	FrequencyPenalty float64        `json:"frequency_penalty"`
	BestOf           int            `json:"best_of"`
	LogitBias        map[string]int `json:"logit_bias,omitempty"`
	User             string         `json:"user"`
}

func CreatePrompParams2(prompt string) (P PromptParams2) {
	P.Model = "text-davinci-003"
	P.Prompt = prompt
	P.MaxTokens = 100
	P.Temperature = 1
	P.TopP = 1
	P.N = 1
	P.Stream = false
	P.Echo = false
	P.PressencePenalty = 0
	P.FrequencyPenalty = 0
	P.BestOf = 1
	return
}

func CreateClient2(APIkey, APIurl string) (C Client2) {
	C.APIkey = APIkey
	C.APIurl = APIurl
	return C
}

type Choice2 struct {
	Text         string
	Index        int
	LogProbs     int
	FinishReason string
}

type Usage2 struct {
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
}

type Response2 struct {
	ID      string
	Object  string
	Created int
	Model   string
	Choices []Choice2
	Usage   Usage2
}

type Client2 struct {
	APIurl     string
	APIkey     string
	HTTPclient http.Client
}

func (C Client2) AskGPTansw2(prompt string) (answer string, err error) {
	P := CreatePrompParams2(prompt)
	var JSONparams []byte
	JSONparams, err = json.Marshal(P)
	fmt.Println("\n string(JSONparams):", string(JSONparams))
	fmt.Println("\n JSONparams after marshal:", JSONparams)
	fmt.Println("\n marshal error:", err)
	req, err := http.NewRequest("POST", C.APIurl, strings.NewReader(string(JSONparams)))
	if err != nil {
		fmt.Println("create request error:", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+C.APIkey)
	fmt.Println("\n req:", req)
	fmt.Println("\n err from new request:", err)
	// Send the request and retrieve the response
	resp, err := C.HTTPclient.Do(req)
	fmt.Println("\n resp:", resp)
	fmt.Println("\n err:", err)
	if err != nil {
		fmt.Println("send request error:", err)
		return
	}
	defer resp.Body.Close()
	// Read and parse the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read and parse response error:", err)
		return
	}
	fmt.Println("\n body:", body)
	fmt.Println("\n string(body):", string(body))
	var response Response2
	json.Unmarshal(body, &response)
	fmt.Println("\n response.Choices[0].Text:", response.Choices[0].Text)
	answer = response.Choices[0].Text
	return answer, err
}

func main() {
	Client2 := CreateClient2(APIkey, "https://api.openai.com/v1/completions")
	answ, err := Client2.AskGPTansw2("what is a ferret")
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

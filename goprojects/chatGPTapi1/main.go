package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// need to be capitalized to be exported?, use tags if not capitalized
// Text         string `json:"text"`

type Choice struct {
	Text         string
	Index        int
	LogProbs     int
	FinishReason string
}

type Usage struct {
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
}

type Response struct {
	ID      string
	Object  string
	Created int
	Model   string
	Choices []Choice
	Usage   Usage
}

type PromptParams struct {
	Model            string         `json:"model"` // required
	Prompt           string         `json:"prompt"`
	Suffix           string         `json:"suffix"`
	MaxTokens        int            `json:"max_tokens"`
	Temperature      float64        `json:"temperature"`
	TopP             float64        `json:"top_p"`
	N                int            `json:"n"`
	Stream           bool           `json:"stream"`
	LogProbs         int            `json:"logprobs"`
	Echo             bool           `json:"echo"`
	Stop             string         `json:"stop"`
	PressencePenalty float64        `json:"presence_penalty"`
	FrequencyPenalty float64        `json:"frequency_penalty"`
	BestOf           float64        `json:"best_of"`
	LogitBias        map[string]int `json:"logit_bias"`
	User             string         `json:"user"`
}

var url = "https://api.openai.com/v1/completions"

/*
1.
*/
func main() {
	var p PromptParams
	p.Model = "text-davinci-003"
	p.MaxTokens = 100
	fmt.Println("p:", p)
	answer, err := p.PromptRequest("What is a duck?")
	fmt.Println("answer:", answer)
	fmt.Println("err:", err)
}

func (p PromptParams) PromptRequest(prompt string) (answer string, err error) {
	/*
		jsonParams, err := json.Marshal(p)
		fmt.Println(jsonParams)
		if err != nil {
			fmt.Println("Marshaling error:", err)
		}
		req, err := http.NewRequest("POST", url, strings.NewReader(string(jsonParams)))
		fmt.Println("req:", req)
		if err != nil {
			fmt.Println("string.NewReader error:", err)
			return answer, err
		}
	*/
	P.Model = "text-davinci-003"
	P.Prompt = "what is a Axolotl"
	var resp2 Response
	resp2, err = askGPT("what is a duck?")
	fmt.Println("resp:", resp2)
	fmt.Println("err:", err)

	question := "What is a duck?"
	req, err := http.NewRequest("POST", url, strings.NewReader(`{
		"model": "text-davinci-003",
		"prompt": "`+question+`",
		"max_tokens": 100
	}`))
	req.Header.Add("Content-Type", "application/json")
	fmt.Println("APIkey:", APIkey)
	fmt.Println("Authorization Header:", "Authorization", "Bearer "+APIkey)
	req.Header.Add("Authorization", "Bearer "+APIkey)
	// Send the request and retrieve the response
	client := &http.Client{}
	var HTTPresp *http.Response
	HTTPresp, err = client.Do(req)
	fmt.Println("HTTPresp:", HTTPresp)
	if err != nil {
		fmt.Println("askGTP() send request error:", err)
		return answer, err
	}
	defer HTTPresp.Body.Close()
	// Read and parse the response
	body, err := ioutil.ReadAll(HTTPresp.Body)
	if err != nil {
		fmt.Println("read and parse response error:", err)
		return answer, err
	}
	fmt.Println("body:", body)
	err = json.Unmarshal(body, &answer)
	var resp Response
	fmt.Println("unmarshaled body:", &resp)
	if err != nil {
		fmt.Println("askGPT() unmarshal error:", err)
		return answer, err
	}
	return answer, err
}

var APIkey = os.Getenv("CHATGPT_API_KEY")

/*
	fmt.Println("askPostHandler running")
	question := c.PostForm("question")
	fmt.Println("question:", question)
	fmt.Println("askGPT running")
	resp, err := askGPT(question)
	if err != nil {
		fmt.Println("error asking chatGPT:", err)
		c.HTML(http.StatusInternalServerError, "index.html", gin.H{"message": "sorry, something went wrong"})
		return
	}
	fmt.Println("c.HTML running")
	// c.HTML(http.StatusAccepted, "index.html", gin.H{"message": resp.Choices[0].Text})
	fmt.Println("resp.Choices[0].Test:", resp.Choices[0].Text)
	c.HTML(http.StatusAccepted, "index.html", gin.H{"message": resp.Choices[0].Text})
*/

// askGPT uses API to ask chatGPT a question and returns response
func askGPT(question string) (resp Response, err error) {
	// Set up the HTTP request, had to change url from suggested to docs example
	// url := "https://api.openai.com/v1/chat/gpt"
	url := "https://api.openai.com/v1/completions"
	fmt.Println(`{
		"model": "text-davinci-003",
		"prompt": "` + question + `",
		"max_tokens": 100
	}`)
	req, err := http.NewRequest("POST", url, strings.NewReader(`{
		"model": "text-davinci-003",
		"prompt": "`+question+`",
		"max_tokens": 100
	}`))
	fmt.Println("req:", req)
	if err != nil {
		fmt.Println("askGPT() create request error:", err)
		return resp, err
	}
	req.Header.Add("Content-Type", "application/json")
	fmt.Println("APIkey:", APIkey)
	fmt.Println("Authorization Header:", "Authorization", "Bearer "+APIkey)
	req.Header.Add("Authorization", "Bearer "+APIkey)
	// Send the request and retrieve the response
	client := &http.Client{}
	var HTTPresp *http.Response
	HTTPresp, err = client.Do(req)
	fmt.Println("HTTPresp:", HTTPresp)
	if err != nil {
		fmt.Println("askGTP() send request error:", err)
		return resp, err
	}
	defer HTTPresp.Body.Close()

	// Read and parse the response
	body, err := ioutil.ReadAll(HTTPresp.Body)
	if err != nil {
		fmt.Println("read and parse response error:", err)
		return resp, err
	}
	fmt.Println("body:", body)
	err = json.Unmarshal(body, &resp)
	fmt.Println("unmarshaled body:", body)
	if err != nil {
		fmt.Println("askGPT() unmarshal error:", err)
		return resp, err
	}
	return resp, err
}

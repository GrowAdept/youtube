package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Set up the HTTP request, had to change url from suggested to docs example
	// url := "https://api.openai.com/v1/chat/gpt"
	url := "https://api.openai.com/v1/completions"
	req, err := http.NewRequest("POST", url, strings.NewReader(`{
		"model": "text-davinci-003",
		"prompt": "what is a cat",
		"max_tokens": 300
	}`))
	if err != nil {
		fmt.Println("create request error:", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	APIkey := os.Getenv("CHATGPT_API_KEY")
	req.Header.Add("Authorization", "Bearer "+APIkey)
	// Send the request and retrieve the response
	client := &http.Client{}
	resp, err := client.Do(req)
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
	fmt.Println("body:", string(body))

	// var response Response
	var response interface{}
	json.Unmarshal(body, &response)
	fmt.Println("response:", response)
	// fmt.Println("response[\"choices\"][0]:", response["choices"][0])
	// choices := response["choices"][0]
	// fmt.Println("choices:", choices)
	// fmt.Println("choices[text]:", choices["text"])
	// fmt.Println("response[\"choices\"][0]:", response["choices"]["text"][0])
}

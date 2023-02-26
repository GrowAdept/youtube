package chatGPTapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type PromptParams struct {
	Model            string         `json:"model"` // required
	Prompt           string         `json:"prompt"`
	Suffix           string         `json:"suffix,omitempty"`
	MaxTokens        int            `json:"max_tokens,omitempty"`
	Temperature      float64        `json:"temperature,omitempty"`
	TopP             float64        `json:"top_p,omitempty"`
	N                int            `json:"n,omitempty"`
	Stream           bool           `json:"stream,omitempty"`
	LogProbs         int            `json:"logprobs,omitempty"`
	Echo             bool           `json:"echo,omitempty"`
	Stop             string         `json:"stop,omitempty"`
	PressencePenalty float64        `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64        `json:"frequency_penalty,omitempty"`
	BestOf           int            `json:"best_of,omitempty"`
	LogitBias        map[string]int `json:"logit_bias,omitempty"`
	User             string         `json:"user,omitempty"`
}

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

type Client struct {
	APIurl     string
	APIkey     string
	HTTPclient http.Client
}

func CreateClient(APIkey, APIurl string) (C Client) {
	C.APIkey = APIkey
	C.APIurl = APIurl
	return C
}

func (C Client) AskGPTansw(prompt string) (answer string, err error) {
	P := CreatePrompParams(prompt)
	var JSONparams []byte
	JSONparams, err = json.Marshal(P)
	fmt.Println("\n JSONparams after marshal:", JSONparams)
	fmt.Println("\n string(JSONparams):", string(JSONparams))
	fmt.Println("\n marshal error:", err)
	req, err := http.NewRequest("POST", C.APIurl, strings.NewReader(string(JSONparams)))
	if err != nil {
		fmt.Println("create request error:", err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+C.APIkey)
	// Send the request and retrieve the response
	resp, err := C.HTTPclient.Do(req)
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
	var response Response
	json.Unmarshal(body, &response)
	fmt.Println("response:", response)
	fmt.Println("\n response.Choices[0].Text:", response.Choices[0].Text)
	answer = response.Choices[0].Text
	return answer, err
}

func (C Client) AskGPTresp(prompt string) (resp *http.Response, err error) {
	fmt.Println("\n askGPTAnsw running")
	P := CreatePrompParams(prompt)
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
	resp, err = C.HTTPclient.Do(req)
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
	var response Response
	json.Unmarshal(body, &response)
	return resp, err
}

func CreatePrompParams(prompt string) (P PromptParams) {
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

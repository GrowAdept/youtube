package main

import (
	"bufio"
	"fmt"
	"os"

	passwordvalidator "github.com/wagslane/go-password-validator"
)

const minEntropyBits = 60

func main() {
	// warning many passwords include profanity or dirty words
	file, err := os.Open("10k-most-common.txt")
	if err != nil {
		fmt.Println("err:", err)
	}
	defer file.Close()

	var entropy float64
	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	fmt.Println("len of words:", len(words))
	for _, v := range words {
		entropy = passwordvalidator.GetEntropy(v)
		if entropy > minEntropyBits {
			fmt.Println("entropy:", entropy)
		}
	}
}

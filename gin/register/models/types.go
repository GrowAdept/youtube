package models

import "fmt"

type User struct {
	ID        string
	Username  string
	Email     string
	pswdHash  string
	CreatedAt string
	Active    string
	verHash   string
	timeout   string
}

func doSomething() {
	fmt.Println("it did something")
}

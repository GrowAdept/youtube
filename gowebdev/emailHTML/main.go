package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
)

func main() {
	// sender data
	from := os.Getenv("FromEmailAddr") //ex: "John.Doe@gmail.com"
	password := os.Getenv("SMTPpwd")   // ex: "ieiemcjdkejspqz"
	// receiver address
	toEmail := os.Getenv("ToEmailAddr") // ex: "Jane.Smith@yahoo.com"
	to := []string{toEmail}
	// smtp - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	// Set up authentication information.
	auth := smtp.PlainAuth("", from, password, host)
	msg := []byte(
		"From: Grow Adept <" + from + ">\r\n" +
			"To: " + toEmail + "\r\n" +
			"Subject: Now with HTML!\r\n" +
			"MIME: MIME-version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
			"\r\n" +
			// "<html><h1>Golang Gophers</h1><ul><li>Robert Griesemer</li><li>Rob Pike</li><li>Ken Thompson</li></ul></html>")
			`<html>
				<h1>Designers of Golang</h1>
				<ul>
					<li>Robert Griesemer</li>
					<li>Rob Pike</li>
					<li>Ken Thompson</li>
				</ul>
			</html>`)
	err := smtp.SendMail(address, auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Check for sent email!")
}

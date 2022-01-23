package main

import (
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
		"From: Justin White <" + from + ">\r\n" +
			"To: " + toEmail + "\r\n" +
			"Subject: Sender Name Test\r\n" +
			"\r\n" +
			"This is the email body.\r\n")
	err := smtp.SendMail(address, auth, from, to, msg)
	if err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"fmt"
	"net/smtp"
)

/*
	WARNING: if using Google to send emails, less secure apps are no longer
	allowed to send emails (5-30-22), app password must be used, app passwords
	are more safe and should be used anyway
*/
func (u *User) SendEmail(subject, HTMLbody string) error {
	// sender data
	to := []string{u.Email}
	// smtp - Simple Mail Transfer Protocol
	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port
	// Set up authentication information.
	auth := smtp.PlainAuth("", fromEmail, SMTPpassword, host)
	msg := []byte(
		"From: " + EntityName + ": <" + fromEmail + ">\r\n" +
			"To: " + u.Email + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME: MIME-version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
			"\r\n" +
			HTMLbody)
	err := smtp.SendMail(address, auth, fromEmail, to, msg)
	if err != nil {
		return err
	}
	fmt.Println("Check for sent email!")
	return nil
}

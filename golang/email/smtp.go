package main

import (
	"log"
	"net/smtp"
)

func main() {
	log.Printf("Start send emil")

	// Set up authentication information.
	auth := smtp.PlainAuth("", "", "", "host")
	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	to := []string{"to@mail.ru"}
	msg := []byte("To: to@mail.ru\r\n" +
		"Subject: Hello there!\r\n" +
		"\r\n" +
		"This is the email body.\r\n")
	_ = auth
	err := smtp.SendMail("host:25", nil, "from@mail.ru", to, msg)
	if err != nil {
		panic(err)
	}
}

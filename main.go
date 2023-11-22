package main

import (
	"fmt"
	"net/smtp"
)

func sendMail() {
	auth := smtp.PlainAuth(
		"",
		"mkbastronaut441@gmail.com",
		"tvfkfzcnvgkungsv",
		"smtp.gmail.com",
	)
	msg := "Subject: Testing"

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"mkbastronaut441@gmail.com",
		[]string{"manishkumar021209@gmail.com"},
		[]byte(msg),
	)

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	sendMail()
}

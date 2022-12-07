package main

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func main() {
	ssl := true
	host := "smtp.126.com"
	username := "chichu_1439@126.com"
	password := "QUQVPDKNDVQUEBNG"
	auth := smtp.PlainAuth("", username, password, host)

	e := email.NewEmail()
	e.From = "chichu_1439@126.com"
	e.To = []string{"zhouchung22@gmail.com"}
	e.Subject = "subject"
	e.Text = []byte("test")
	var err error
	if ssl {
		err = e.SendWithTLS("smtp.126.com:465", auth, &tls.Config{ServerName: host})
	} else {
		err = e.Send("smtp.126.com:25", auth)
	}
	if err != nil {
		fmt.Printf("format string %v", err)
	}
}

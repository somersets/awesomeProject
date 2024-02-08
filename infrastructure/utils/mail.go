package utils

import (
	"fmt"
	"net/smtp"
)

type MailService interface {
	SendMail(message string)
}

type mailConfig struct {
	emailFrom    string
	passwordFrom string
	emailTo      []string
	smtpHost     string
	smtpPort     string
	identity     string
}

func NewMailService(identity string, emailFrom string, passwordFrom string, emailTo string) MailService {
	from := emailFrom
	password := passwordFrom

	// Receiver email address.
	to := []string{
		emailTo,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	return &mailConfig{
		emailFrom:    from,
		passwordFrom: password,
		emailTo:      to,
		smtpHost:     smtpHost,
		smtpPort:     smtpPort,
		identity:     identity,
	}

}

func (mc *mailConfig) SendMail(message string) {
	auth := smtp.PlainAuth(mc.identity, mc.emailFrom, mc.passwordFrom, mc.smtpHost)
	err := smtp.SendMail(mc.smtpHost+":"+mc.smtpPort, auth, mc.emailFrom, mc.emailTo, []byte(message))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}

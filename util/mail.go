package util

import (
	"eCommerce/config"
	"github.com/caarlos0/env/v6"
	"gopkg.in/mail.v2"
)

var (
	mailConfig = config.MailConfig{}
	_          = env.Parse(&mailConfig)
)

func SendMail(email string, subject string, content string) (err error) {
	m := mail.NewMessage()
	m.SetAddressHeader("From", "no-reply@bipbip.com", "BipBip")
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)
	d := mail.NewDialer(mailConfig.SmtpHost, mailConfig.SmtpPort, mailConfig.SmtpUser, mailConfig.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

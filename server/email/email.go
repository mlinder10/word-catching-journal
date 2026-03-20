package email

import (
	"net/smtp"

	"github.com/mlinder10/wcj/config"
)

func Send(to, subject, body string) error {
	from := "Subject: " + subject + "\r\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	msg := []byte(from + mime + body)

	return smtp.SendMail(
		config.Env.EmailAddr,
		smtp.PlainAuth(
			config.Env.EmailIdentity,
			config.Env.EmailAddress,
			config.Env.EmailPassword,
			config.Env.EmailHost,
		),
		config.Env.EmailAddress,
		[]string{to},
		[]byte(msg),
	)
}

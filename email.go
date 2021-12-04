package email

import (
	"fmt"
	"github.com/czasg/go-fill"
	"net/smtp"
	"strings"
)

type Payload struct {
	UserAccount string `env:"EMAIL_ACCOUNT"`
	Secret      string `env:"EMAIL_SECRET"`
	Subject     string `default:"email by go"`
	To          []string
	Body        string `default:"email by go"`
	Host        string `default:"smtp.qq.com"`
	Port        string `default:"25"`
	ContentType string `default:"text/plain; charset=\"utf-8\""`
	MIMEVersion string `default:"1.0"`
}

func SendMail(payload Payload) error {
	if len(payload.To) < 1 {
		return nil
	}
	err := fill.Fill(&payload, fill.OptEnv, fill.OptDefault)
	if err != nil {
		return err
	}
	message := []string{
		fmt.Sprintf("Subject: %s", payload.Subject),
		fmt.Sprintf("Content-Type: %s", payload.ContentType),
		fmt.Sprintf("MIME-Version: %s", payload.MIMEVersion),
		fmt.Sprintf("From: %s", payload.UserAccount),
		fmt.Sprintf("\r\n%s", payload.Body),
	}
	auth := smtp.PlainAuth("", payload.UserAccount, payload.Secret, payload.Host)
	return smtp.SendMail(
		payload.Host+":"+payload.Port,
		auth,
		payload.UserAccount,
		payload.To,
		[]byte(strings.Join(message, "\r\n")),
	)
}

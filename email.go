package email

import (
	"fmt"
	"github.com/czasg/go-fill"
	"net/smtp"
	"strings"
)

type Payload struct {
	UserAccount string   `json:"userAccount" env:"EMAIL_ACCOUNT"`
	Secret      string   `json:"secret" env:"EMAIL_SECRET"`
	Subject     string   `json:"subject" default:"email by go"`
	To          []string `json:"to"`
	Body        string   `json:"body" default:"email by go"`
	Host        string   `json:"host" default:"smtp.qq.com"`
	Port        string   `json:"port" default:"25"`
	ContentType string   `json:"contentType" default:"text/plain; charset=\"utf-8\""`
	MIMEVersion string   `json:"mimeVersion" default:"1.0"`
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

package email

import (
	"net/smtp"
	"os"
	"strings"
)

type Email struct {
	From     string
	Password string
	Host     string
	Port     string
}

func NewEmail(from, password, host, port string) *Email {
	return &Email{
		From:     from,
		Password: password,
		Host:     host,
		Port:     port,
	}
}

type EmailTemplate string

const (
	emailHeader string = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\""

	EmailTemplateUserCreated EmailTemplate = "../pkg/email/template/user_created.html"
)

// Send sends an email to the specified recipient with the specified subject and body.
func (e *Email) Send(to []string, subject string, template EmailTemplate, templateVariables map[string]string) error {
	auth := smtp.PlainAuth("", e.From, e.Password, e.Host)

	basePath, err := os.Getwd()
	if err != nil {
		return err
	}

	file, err := os.ReadFile(basePath + "/" + string(template))
	if err != nil {
		return err
	}

	body := string(file)

	for key, value := range templateVariables {
		body = strings.ReplaceAll(body, key, value)
	}

	return smtp.SendMail(e.Host+":"+e.Port, auth, e.From, to, []byte("Subject: "+subject+"\n"+emailHeader+"\n\n"+body))
}

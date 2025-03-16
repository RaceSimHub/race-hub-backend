package email

import (
	"net/smtp"
	"net/url"
	"os"
	"strings"
)

type Email struct {
	ApplicationHost string
	From            string
	Password        string
	Host            string
	Port            string
}

func NewEmail(applicationHost, from, password, host, port string) *Email {
	return &Email{
		ApplicationHost: applicationHost,
		From:            from,
		Password:        password,
		Host:            host,
		Port:            port,
	}
}

type EmailTemplate string

const (
	emailHeader string = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\""

	EmailTemplateUserCreated    EmailTemplate = "../pkg/email/template/user_created.html"
	EmailTemplateForgotPassword EmailTemplate = "../pkg/email/template/forgot_password.html"
)

func (e *Email) SendUserCreatedEmail(email, name, token string) error {
	return e.send([]string{email}, "Confirme seu email", EmailTemplateUserCreated, map[string]string{
		"{NAME}":               name,
		"{CONFIRMATION_TOKEN}": token,
	})
}

func (e *Email) SendForgotPasswordEmail(email, token string) error {
	url := e.ApplicationHost + "/reset-password?email=" + url.QueryEscape(email) + "&token=" + token

	return e.send([]string{email}, "Recuperação de senha", EmailTemplateForgotPassword, map[string]string{
		"{URL_RESET_PASSWORD}": url,
	})
}

func (e *Email) send(to []string, subject string, template EmailTemplate, templateVariables map[string]string) error {
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

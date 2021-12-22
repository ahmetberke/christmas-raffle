package mailer

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"

	"github.com/ahmetberke/christmas-raffle/configs"
)

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

type Mail struct {
	to      []string
	subject string
	body    string
}

func NewMail(to []string, subject string) *Mail {
	return &Mail{
		to:      to,
		subject: subject,
	}
}

// Parse template with data
func (m *Mail) parseTemplate(templateDir string, data interface{}) error {
	// parse html file
	t, err := template.ParseFiles(templateDir)
	if err != nil {
		return err
	}
	// implement data with buffer
	buffer := new(bytes.Buffer)
	if err = t.Execute(buffer, data); err != nil {
		return err
	}
	m.body = buffer.String()
	return nil
}

func (m *Mail) sendMail() bool {
	body := fmt.Sprintf("To: %v\r\nSubject: %v\r\n%v\n\r%v", m.to[0], m.subject, MIME, m.body)
	SMTP := fmt.Sprintf("%s:%s", configs.Manager.SmtpSettings.Host, configs.Manager.SmtpSettings.Port)
	if err := smtp.SendMail(SMTP, smtp.PlainAuth("", configs.Manager.SmtpSettings.Email, configs.Manager.SmtpSettings.Password, configs.Manager.SmtpSettings.Host), configs.Manager.SmtpSettings.Email, m.to, []byte(body)); err != nil {
		return false
	}
	return true
}

func (m *Mail) Send(templateDir string, data interface{}) {
	err := m.parseTemplate(templateDir, data)
	if err != nil {
		log.Fatal(err)
	}
	if ok := m.sendMail(); ok {
		log.Printf("Email has been sent to %s\n", m.to)
	} else {
		log.Printf("Failed to send the email to %s\n", m.to)
	}
}

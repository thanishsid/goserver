package mailer

import (
	"bytes"
	"context"
	"text/template"

	gomail "gopkg.in/mail.v2"
)

type mailer struct {
	Dialer    *gomail.Dialer
	Templates *template.Template
}

type mailData struct {
	templateName string
	templateData any
	to           []string
	subject      string
}

// Send email with data and specific template.
func sendMail(ctx context.Context, m *mailer, data mailData) error {
	templateBuffer := new(bytes.Buffer)

	if err := m.Templates.ExecuteTemplate(
		templateBuffer,
		data.templateName,
		data.templateData); err != nil {
		return err
	}

	msg := gomail.NewMessage()

	msg.SetHeader("From", m.Dialer.Username)
	msg.SetHeader("To", data.to...)
	msg.SetHeader("Subject", data.subject)
	msg.SetBody("text/html", templateBuffer.String())

	return m.Dialer.DialAndSend(msg)
}

// Send Link Mail.
func (m *mailer) SendLinkMail(ctx context.Context, data LinkMailData) error {
	return sendMail(ctx, m, mailData{
		templateName: "link-mail.html",
		to:           []string{data.To},
		subject:      data.Subject,
		templateData: NewDataWithDefaults(data),
	})
}

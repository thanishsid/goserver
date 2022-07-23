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
	TemplateName string
	TemplateData any
	To           []string
	Subject      string
}

// Send email with data and specific template.
func (m *mailer) sendMail(ctx context.Context, data mailData) error {
	templateBuffer := new(bytes.Buffer)

	if err := m.Templates.ExecuteTemplate(
		templateBuffer,
		data.TemplateName,
		data.TemplateData); err != nil {
		return err
	}

	msg := gomail.NewMessage()

	msg.SetHeader("From", m.Dialer.Username)
	msg.SetHeader("To", data.To...)
	msg.SetHeader("Subject", data.Subject)
	msg.SetBody("text/html", templateBuffer.String())

	return m.Dialer.DialAndSend(msg)
}

// Send Link Mail.
func (m *mailer) SendLinkMail(ctx context.Context, data LinkMailTemplateData) error {
	return m.sendMail(ctx, mailData{
		TemplateName: "link-mail.html",
		TemplateData: NewDataWithDefault(data),
	})
}

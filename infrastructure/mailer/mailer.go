package mailer

import (
	"context"
	"crypto/tls"
	"io/fs"
	"text/template"
	"time"

	"github.com/thanishsid/goserver/config"
	gomail "gopkg.in/mail.v2"
)

type Mailer interface {
	SendLinkMail(ctx context.Context, data LinkMailData) error
}

type MailerConfig struct {
	DialerTimeout   time.Duration
	DialerTLSConfig *tls.Config
	TemplateStore   fs.FS
	TemplatePaths   []string
}

func NewMailer(conf MailerConfig) (Mailer, error) {
	dialer := gomail.NewDialer(
		config.C.MailHost,
		config.C.MailPort,
		config.C.MailUser,
		config.C.MailPassword,
	)
	dialer.TLSConfig = conf.DialerTLSConfig
	dialer.Timeout = conf.DialerTimeout

	sc, err := dialer.Dial()
	if err != nil {
		return nil, err
	}

	err = sc.Close()
	if err != nil {
		return nil, err
	}

	templates, err := template.ParseFS(conf.TemplateStore, conf.TemplatePaths...)
	if err != nil {
		return nil, err
	}

	return &mailer{
		Dialer:    dialer,
		Templates: templates,
	}, nil
}

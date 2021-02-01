package mail

import (
	"net/smtp"

	"github.com/jordan-wright/email"
	"github.com/sirupsen/logrus"
)

// Config ...
type Config struct {
	SMTPAddr string
	Identity string
	UserName string
	Password string
	Host     string
}

// Send ...
func Send(config Config, from, to, subject, test, html string) error {
	e := email.NewEmail()
	e.From = from
	e.To = []string{to}
	e.Subject = subject
	e.Text = []byte(test)
	e.HTML = []byte(html)
	err := e.Send(config.SMTPAddr, smtp.PlainAuth(config.Identity, config.UserName, config.Password, config.Host))
	if err != nil {
		logrus.WithError(err).Error()
	}
	return err
}

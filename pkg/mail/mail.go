package mail

import (
	"hanya-go/pkg/config"
	"hanya-go/pkg/logger"
	"sync"
)

type Email struct {
	FromAddress string
	FromName    string
	To          []string
	Bcc         []string
	Cc          []string
	Subject     string
	Text        []byte
	HTML        []byte
	Tls         bool
}

type Mailer struct {
	Driver Driver
}

var once sync.Once
var internalMailer *Mailer
var mode string

func NewMailer() *Mailer {
	once.Do(func() {
		mode = config.Get("mail.default")

		if mode == "smtp" {
			internalMailer = &Mailer{
				Driver: &SMTP{},
			}
		} else if mode == "sendcloud" {
			internalMailer = &Mailer{
				Driver: &SENDCLOUD{},
			}
		}
	})
	return internalMailer
}

func (mailer *Mailer) Send(email Email) bool {

	logger.DebugString("mail/mail.go 47", "mode", mode)

	if mode == "smtp" {
		logger.DebugString("mail/mail.go 47", "stmp", mode)
		return mailer.Driver.Send(email, config.GetStringMapString("mail.smtp"))
	} else if mode == "sendcloud" {
		return mailer.Driver.Send(email, config.GetStringMapString("mail.sendcloud"))
	} else {
		return false
	}
}

package mail

import (
	"hanya-go/pkg/config"
	"sync"
)

type From struct {
	Address string
	Name    string
}

type Email struct {
	From    From
	To      []string
	Bcc     []string
	Cc      []string
	Subject string
	Text    []byte
	HTML    []byte
}

type Mailer struct {
	Driver Driver
}

var once sync.Once
var internalMailer *Mailer

var mode string = config.GetString("sms.default")

func NewMailer() *Mailer {
	once.Do(func() {

		if mode == "smtp" {
			internalMailer = &Mailer{
				Driver: &SMTP{},
			}
		} else if mode == "sendcloud" {
			// TODO SENDCLOUD
		}
	})
	return internalMailer
}

func (mailer *Mailer) Send(email Email) bool {
	if mode == "smtp" {
		return mailer.Driver.Send(email, config.GetStringMapString("mail.smtp"))
	} else if mode == "sendcloud" {
		// TODO SENDCLOUD
	} else {
		return false
	}

	return false
}

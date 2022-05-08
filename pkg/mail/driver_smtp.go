package mail

import (
	"crypto/tls"
	"fmt"
	emailPKG "github.com/jordan-wright/email"
	"hanya-go/pkg/logger"
	"net/smtp"
)

// SMTP 实现 email.Driver interface
type SMTP struct{}

// Send 实现 email.Driver interface 的 Send 方法
func (s *SMTP) Send(email Email, config map[string]string) bool {
	e := emailPKG.NewEmail()

	logger.DebugJSON("SMTP发送邮件", "EMAIL对象原始信息", email)

	e.From = fmt.Sprintf("%v <%v>", email.FromName, email.FromAddress)
	e.To = email.To
	e.Bcc = email.Bcc
	e.Cc = email.Cc
	e.Subject = email.Subject
	e.Text = email.Text
	e.HTML = email.HTML

	logger.DebugJSON("SMTP发送邮件", "发件详情", e)

	var err error
	// 判断是否需要用TLS连接
	if email.Tls {
		err = e.SendWithTLS(
			fmt.Sprintf("%v:%v", config["host"], config["port"]),

			smtp.PlainAuth(
				"",
				config["username"],
				config["password"],
				config["host"],
			),
			&tls.Config{
				ServerName: config["host"],
			},
		)
	} else {
		err = e.Send(
			fmt.Sprintf("%v:%v", config["host"], config["port"]),

			smtp.PlainAuth(
				"",
				config["username"],
				config["password"],
				config["host"],
			),
		)
	}

	if err != nil {
		logger.ErrorString("SMTP发送邮件", "发送出错", err.Error())
		return false
	}

	logger.DebugString("发送邮件", "发件成功", "")

	return true

}

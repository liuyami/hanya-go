// Package config 站点配置信息
package config

import "hanya-go/pkg/config"

func init() {
	config.Add("mail", func() map[string]interface{} {
		return map[string]interface{}{

			"default": "smtp",

			// 默认是 Mailhog 的配置
			"smtp": map[string]interface{}{
				"host":     config.Env("MAIL_SMTP_HOST", ""),
				"port":     config.Env("MAIL_SMTP_PORT"),
				"username": config.Env("MAIL_SMTP_USERNAME", ""),
				"password": config.Env("MAIL_SMTP_PASSWORD", ""),
				"from":     config.Env("MAIL_SMTP_FROM_ADDRESS", ""),
				"name":     config.Env("MAIL_SMTP_FROM_NAME", ""),
			},

			//"from": map[string]interface{}{
			//	"address": config.Env("MAIL_FROM_ADDRESS", "gohub@example.com"),
			//	"name":    config.Env("MAIL_FROM_NAME", "Gohub"),
			//},

			"sendcloud": map[string]interface{}{
				"apiUser": config.Env("MAIL_SENDCLOUD_API_USER", ""),
				"apiKey":  config.Env("MAIL_SENDCLOUD_API_KEY", ""),
				"from":    config.Env("MAIL_SENDCLOUD_FROM_ADDRESS", ""),
				"name":    config.Env("MAIL_SENDCLOUD_FROM_NAME", ""),
			},
		}
	})
}

// Package config 站点配置信息
package config

import "hanya-go/pkg/config"

func init() {
	config.Add("mail", func() map[string]interface{} {
		return map[string]interface{}{

			"default": "sendcloud",

			// 默认是 STMP，第三方包 Mailhog 的配置
			"smtp": map[string]interface{}{
				"host":         config.Env("MAIL_SMTP_HOST", ""),
				"port":         config.Env("MAIL_SMTP_PORT", ""),
				"username":     config.Env("MAIL_SMTP_USERNAME", ""),
				"password":     config.Env("MAIL_SMTP_PASSWORD", ""),
				"tls":          config.Env("MAIL_SMTP_NEED_TLS", false),
				"from_address": config.Env("MAIL_SMTP_FROM_ADDRESS", ""),
				"from_name":    config.Env("MAIL_SMTP_FROM_NAME", ""),
			},

			"sendcloud": map[string]interface{}{
				"apiUser":      config.Env("MAIL_SENDCLOUD_API_USER", ""),
				"apiKey":       config.Env("MAIL_SENDCLOUD_API_KEY", ""),
				"from_address": config.Env("MAIL_SENDCLOUD_FROM_ADDRESS", ""),
				"from_name":    config.Env("MAIL_SENDCLOUD_FROM_NAME", ""),
			},
		}
	})
}

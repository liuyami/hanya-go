// Package config 站点配置信息
package config

import "hanya-go/pkg/config"

func init() {
	config.Add("sms", func() map[string]interface{} {
		return map[string]interface{}{

			// 默认是阿里云的测试 sign_name 和 template_code
			"aliyun": map[string]interface{}{
				"access_key_id":     config.Env("SMS_ALIYUN_ACCESS_ID", ""),
				"access_key_secret": config.Env("SMS_ALIYUN_ACCESS_SECRET", ""),
				"sign_name":         config.Env("SMS_ALIYUN_SIGN_NAME", ""),
				"template_code":     config.Env("SMS_ALIYUN_TEMPLATE_CODE", ""),
			},
		}
	})
}
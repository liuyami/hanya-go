package verifycode

import (
	"fmt"
	"hanya-go/pkg/app"
	"hanya-go/pkg/config"
	"hanya-go/pkg/helpers"
	"hanya-go/pkg/logger"
	"hanya-go/pkg/mail"
	"hanya-go/pkg/redis"
	"hanya-go/pkg/sms"
	"strings"
	"sync"
)

type VerifyCode struct {
	Store Store
}

var once sync.Once
var internalVerifyCode *VerifyCode

func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				KeyPrefix:   config.GetString("app.name") + ":verifycode:",
			},
		}
	})

	return internalVerifyCode
}

// SendSMS 发送短信验证码，调用示例：
//         verifycode.NewVerifyCode().SendSMS(request.Phone)
func (vc *VerifyCode) SendSMS(phone string) bool {
	// 生成验证码
	code := vc.generateVerifyCode(phone)

	// 方便本地和 API 自动测试
	if !app.IsProduction() && strings.HasPrefix(phone, config.GetString("verifycode.debug_phone_prefix")) {
		return true
	}

	// 发送信息
	return sms.NewSMS().Send(phone, sms.Message{
		Template: config.GetString("sms.aliyun.template_code"),
		Data:     map[string]string{"code": code},
	})
}

func (vc *VerifyCode) SendEmail(email string) error {
	// 生成验证码
	code := vc.generateVerifyCode(email)

	// 方便本地和 API 自动测试
	if !app.IsProduction() && strings.HasPrefix(email, config.GetString("verifycode.debug_email_suffix")) {
		return nil
	}

	content := fmt.Sprintf("<h1>您的 Email 验证码是 %v </h1>", code)

	// 发件信息
	mailMode := config.Get("mail.default")
	fromAddress := ""
	fromName := ""

	if mailMode == "smtp" {
		fromName = config.GetString("mail.smtp.from_name")
		fromAddress = config.GetString("mail.smtp.from_address")
	} else if mailMode == "sendcloud" {
		fromName = config.GetString("mail.sendcloud.from_name")
		fromAddress = config.GetString("mail.sendcloud.from_address")
	}

	// 发送邮件
	res := mail.NewMailer().Send(mail.Email{
		FromName:    fromName,
		FromAddress: fromAddress,
		To:          []string{email},
		Subject:     "验证码",
		HTML:        []byte(content),
		Tls:         config.GetBool("mail.smtp.tls"),
	})

	if res {
		logger.DebugString("邮件验证码", "成功", "成功")
	} else {
		logger.DebugString("邮件验证码", "失败", "失败")
	}

	return nil
}

// generateVerifyCode 生成验证码，并放置于 Redis 中
func (vc *VerifyCode) generateVerifyCode(key string) string {

	// 生成随机码
	code := helpers.RandomNumber(config.GetInt("verifycode.code_length"))

	// 为方便开发，本地环境使用固定验证码
	if app.IsLocal() {
		code = config.GetString("verifycode.debug_code")
	}

	logger.DebugJSON("验证码", "生成验证码", map[string]string{key: code})

	vc.Store.Set(key, code)

	return code
}

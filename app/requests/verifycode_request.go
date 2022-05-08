package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"hanya-go/app/requests/validators"
)

type VerifyCodePhoneRequest struct {
	CaptchaID     string `json:"captcha_id" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer" valid:"captcha_answer"`

	Phone string `json:"phone,omitempty" valid:"phone"`
}

type VerifyCodeEmailRequest struct {
	CaptchaID     string `json:"captcha_id" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer" valid:"captcha_answer"`

	Email string `json:"email,omitempty" valid:"email"`
}

// VerifyCodePhone 验证表单，返回长度等于零即通过
func VerifyCodePhone(data interface{}, c *gin.Context) map[string][]string {

	// 1. 定义规则
	rules := govalidator.MapData{
		"phone":          []string{"required", "digits:11"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:4"},
	}

	// 2. 定义错误信息
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号码必须填写",
			"digits:手机号码必须为11位数字",
		},
		"captcha_id": []string{
			"required:图片验证码必须填写",
		},
		"captcha_answer": []string{
			"required:图片验证码答案必填",
			"digits:图片验证码长度必须为 4 位的数字",
		},
	}

	errs := validate(data, rules, messages)

	// 图形验证码
	_data := data.(*VerifyCodePhoneRequest)

	errs = validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)

	return errs
}

// VerifyCodeEmail  验证表单，返回长度等于零即通过
func VerifyCodeEmail(data interface{}, c *gin.Context) map[string][]string {

	// 1. 定义规则
	rules := govalidator.MapData{
		"email":          []string{"required", "min:4", "max:64", "email"},
		"captcha_id":     []string{"required"},
		"captcha_answer": []string{"required", "digits:4"},
	}

	// 2. 定义错误信息
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
		"captcha_id": []string{
			"required:图片验证码必须填写",
		},
		"captcha_answer": []string{
			"required:图片验证码答案必填",
			"digits:图片验证码长度必须为 4 位的数字",
		},
	}

	errs := validate(data, rules, messages)

	// 图形验证码
	_data := data.(*VerifyCodeEmailRequest)

	errs = validators.ValidateCaptcha(_data.CaptchaID, _data.CaptchaAnswer, errs)

	return errs
}

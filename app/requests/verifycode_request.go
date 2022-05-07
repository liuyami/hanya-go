package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"hanya-go/pkg/captcha"
)

type VerifyCodePhoneRequest struct {
	CaptchaID     string `json:"captcha_id" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer" valid:"captcha_answer"`

	Phone string `json:"phone,omitempty" valid:"phone"`
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

	if ok := captcha.NewCaptcha().VerifyCaptcha(_data.CaptchaID, _data.CaptchaAnswer); !ok {
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}

	return errs
}

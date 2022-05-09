package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"hanya-go/app/requests/validators"
)

// SignUpPhoneExistRequest 手机是否已使用请求结构
type SignUpPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

// SignUpPhoneExist 检测手机号码是否已使用
func SignUpPhoneExist(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}
	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号码必须填写",
			"digits:手机号码长度必须为11位数字",
		},
	}

	return validate(data, rules, messages)
}

// SignupEmailExistRequest 邮件地址是否已使用结构体
type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

// SignupEmailExist 检测邮件地址是否已使用
func SignupEmailExist(data interface{}, c *gin.Context) map[string][]string {

	// 自定义验证规则
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}

	return validate(data, rules, messages)
}

// SignupUsingPhoneRequest 通过手机注册的请求体
type SignupUsingPhoneRequest struct {
	Phone           string `json:"phone,omitempty" valid:"phone"`
	VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
	Nickname        string `json:"nickname" valid:"nickname"`
	Password        string `json:"password,omitempty" valid:"password"`
	PasswordConfirm string `json:"password_confirm,omitempty" valid:"password_confirm"`
}

// SignupUsingPhone 手机注册验证
func SignupUsingPhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":            []string{"required", "digits:11", "not_exists:users,phone"},
		"nickname":         []string{"required", "alpha_num", "between:3,20", "not_exists:users,nickname"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
		"verify_code":      []string{"required", "digits:4"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号码必须填写",
			"digits: 手机号码只能11位数字",
		},
		"nickname": []string{
			"required:昵称必须填写",
			"alpha_num: 用户名格式错误，只允许数字和英文",
			"between: 用户名长度需在 3~20 之间",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"password_confirm": []string{
			"required:确认密码框为必填项",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 4 位的数字",
		},
	}

	errs := validate(data, rules, messages)

	_data := data.(*SignupUsingPhoneRequest)

	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}

// SignupUsingEmailRequest   通过电子邮件地址注册的请求体
type SignupUsingEmailRequest struct {
	Email           string `json:"email,omitempty" valid:"email"`
	VerifyCode      string `json:"verify_code,omitempty" valid:"verify_code"`
	Nickname        string `json:"nickname" valid:"nickname"`
	Password        string `json:"password,omitempty" valid:"password"`
	PasswordConfirm string `json:"password_confirm,omitempty" valid:"password_confirm"`
}

// SignupUsingEmail  电子邮件注册验证
func SignupUsingEmail(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"email":            []string{"required", "min:4", "max:30", "email", "not_exists:users,email"},
		"nickname":         []string{"required", "alpha_num", "between:3,20", "not_exists:users,nickname"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
		"verify_code":      []string{"required", "digits:4"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
			"not_exists:Email 已被占用",
		},
		"nickname": []string{
			"required:昵称必须填写",
			"alpha_num: 用户名格式错误，只允许数字和英文",
			"between: 用户名长度需在 3~20 之间",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"password_confirm": []string{
			"required:确认密码框为必填项",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 4 位的数字",
		},
	}

	errs := validate(data, rules, messages)

	_data := data.(*SignupUsingEmailRequest)

	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	errs = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)

	return errs
}

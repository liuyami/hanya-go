package auth

import (
	"errors"
	"hanya-go/app/models/user"
)

// Attempt 尝试登录
func Attempt(email, password string) (user.User, error) {
	userModel := user.GetByEmail(email)

	if userModel.UserID == 0 {
		return user.User{}, errors.New("账号不存在")
	}

	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}

	return userModel, nil
}

// LoginByPhone 手机号码登录
// 注意，手机号码是用验证码，所以不用匹配密码从而和Attempt 区别开来
func LoginByPhone(phone string) (user.User, error) {
	userModel := user.GetByPhone(phone)

	if userModel.UserID == 0 {
		return user.User{}, errors.New("账号不存在")
	}
	return userModel, nil
}

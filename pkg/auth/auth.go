package auth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"hanya-go/app/models/user"
	"hanya-go/pkg/logger"
)

// Attempt 尝试登录
func Attempt(loginID, password string) (user.User, error) {
	userModel := user.GetByMulti(loginID)

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

func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}

	return userModel
}

// CurrentUID 从 gin.context 中获取当前登录用户 ID
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}

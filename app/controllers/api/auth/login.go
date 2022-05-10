package auth

import (
	"github.com/gin-gonic/gin"
	"hanya-go/app/requests"
	"hanya-go/app/response"
	"hanya-go/pkg/auth"
	"hanya-go/pkg/jwt"
)

func LoginByPhone(c *gin.Context) {

	// 1.验证表单
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}

	// 2.登录
	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		response.Fail(c, 1001, "登录失败", err)
	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID())

		response.Success(c, gin.H{
			"token": token,
		})
	}
}

// LoginByPassword 密码登录
func LoginByPassword(c *gin.Context) {

	// 1.验证表单
	req := requests.LoginByPasswordRequest{}
	if ok := requests.Validate(c, &req, requests.LoginByPassword); !ok {
		return
	}

	// 2. 尝试登录
	user, err := auth.Attempt(req.LoginID, req.Password)

	if err != nil {
		response.Fail(c, 1001, "登录失败，错误的账号或密码", nil)
	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID())

		response.Success(c, gin.H{
			"token": token,
		})
	}
}

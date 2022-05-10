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

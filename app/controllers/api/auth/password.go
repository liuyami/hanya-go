package auth

import (
	"github.com/gin-gonic/gin"
	"hanya-go/app/models/user"
	"hanya-go/app/requests"
	"hanya-go/app/response"
)

func ResetPasswordByPhone(c *gin.Context) {

	// 1.表单验证
	req := requests.ResetByPhoneRequest{}
	if ok := requests.Validate(c, &req, requests.ResetByPhone); !ok {
		return
	}

	// 2.更新密码
	userModel := user.GetByPhone(req.Phone)

	if userModel.UserID == 0 {
		response.Fail(c, 1001, "账户异常：不存在或已被禁用", nil)
	} else {
		userModel.Password = req.Password
		userModel.Save()

		response.Success(c, nil)
	}
}

func ResetPasswordByEmail(c *gin.Context) {
	// 1. 验证表单
	request := requests.ResetByEmailRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByEmail); !ok {
		return
	}

	// 2. 更新密码
	userModel := user.GetByEmail(request.Email)
	if userModel.UserID == 0 {
		response.Fail(c, 1001, "账户异常：不存在或已被禁用", nil)
	} else {
		userModel.Password = request.Password
		userModel.Save()
		response.Success(c, nil)
	}
}

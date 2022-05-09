package auth

import (
	"github.com/gin-gonic/gin"
	"hanya-go/app/models/user"
	"hanya-go/app/requests"
	"hanya-go/app/response"
)

//type Signup struct {
//	api.BaseAPI
//}

func IsPhoneExist(c *gin.Context) {

	request := requests.SignUpPhoneExistRequest{}

	if ok := requests.Validate(c, &request, requests.SignUpPhoneExist); !ok {
		return
	}

	response.Success(c, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

func IsEmailExist(c *gin.Context) {

	// 获取请求参数，并做表单验证
	request := requests.SignupEmailExistRequest{}

	if ok := requests.Validate(c, &request, requests.SignupEmailExist); !ok {
		return
	}

	response.Success(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}

// SignupUsingPhone 使用手机和验证码进行注册
func SignupUsingPhone(c *gin.Context) {

	// 1. 验证表单
	request := requests.SignupUsingPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
		return
	}

	// 2. 验证成功，创建数据
	userModel := user.User{
		Nickname: request.Nickname,
		Phone:    request.Phone,
		Password: request.Password,
	}
	userModel.Create()

	if userModel.UserID > 0 {
		response.Success(c, userModel)
	} else {
		response.Fail(c, 10001, "创建用户失败，请稍后尝试~", "")
	}
}

// SignupUsingEmail  使用电子邮件进行注册
func SignupUsingEmail(c *gin.Context) {

	// 1. 验证表单
	request := requests.SignupUsingEmailRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingEmail); !ok {
		return
	}

	// 2. 验证成功，创建数据
	userModel := user.User{
		Nickname: request.Nickname,
		Email:    request.Email,
		Password: request.Password,
	}
	userModel.Create()

	if userModel.UserID > 0 {
		response.Success(c, userModel)
	} else {
		response.Fail(c, 10001, "创建用户失败，请稍后尝试~", "")
	}
}

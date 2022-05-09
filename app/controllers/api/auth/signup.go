package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
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
	_user := user.User{
		Nickname: request.Nickname,
		Phone:    request.Phone,
		Password: request.Password,
	}
	_user.Create()

	if cast.ToUint(_user.ID) > 0 {
		response.Success(c, _user)
	} else {
		response.Fail(c, 10001, "创建用户失败，请稍后尝试~", "")
	}
}

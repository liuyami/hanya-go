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

type isResult struct {
	Exist bool `json:"exist"`
}

func IsPhoneExist(c *gin.Context) {

	request := requests.SignUpPhoneExistRequest{}

	if ok := requests.Validate(c, &request, requests.SignUpPhoneExist); !ok {
		return
	}

	response.Success(c, isResult{
		Exist: user.IsPhoneExist(request.Phone),
	})
}

func IsEmailExist(c *gin.Context) {

	// 获取请求参数，并做表单验证
	request := requests.SignupEmailExistRequest{}

	if ok := requests.Validate(c, &request, requests.SignupEmailExist); !ok {
		return
	}

	response.Success(c, isResult{
		Exist: user.IsPhoneExist(request.Email),
	})
}

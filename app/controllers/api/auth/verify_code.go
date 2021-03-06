package auth

import (
	"github.com/gin-gonic/gin"
	"hanya-go/app/requests"
	"hanya-go/app/response"
	"hanya-go/pkg/captcha"
	"hanya-go/pkg/logger"
	"hanya-go/pkg/verifycode"
)

func GetCaptcha(c *gin.Context) {

	id, base64, err := captcha.NewCaptcha().GenerateCaptcha()

	logger.LogIf(err)

	response.Success(c, gin.H{
		"captcha_id":    id,
		"captcha_image": base64,
	})
}

// SendUsingPhone 发送手机验证码
func SendUsingPhone(c *gin.Context) {

	// 验证请求
	request := requests.VerifyCodePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodePhone); !ok {
		return
	}

	// 发送短信
	if ok := verifycode.NewVerifyCode().SendSMS(request.Phone); !ok {
		response.Fail(c, 500, "发送短信失败", nil)
	} else {
		response.Success(c, "发送成功")
	}

}

func SendUsingEmail(c *gin.Context) {

	// 验证请求
	request := requests.VerifyCodeEmailRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodeEmail); !ok {
		return
	}

	// 发送邮件
	err := verifycode.NewVerifyCode().SendEmail(request.Email)

	if err != nil {
		response.Fail(c, 500, "发送失败", err.Error())
	} else {
		response.Success(c, "")
	}

}

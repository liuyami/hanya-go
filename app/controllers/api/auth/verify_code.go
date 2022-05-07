package auth

import (
	"github.com/gin-gonic/gin"
	"hanya-go/app/response"
	"hanya-go/pkg/captcha"
	"hanya-go/pkg/logger"
)

func GetCaptcha(c *gin.Context) {

	id, base64, err := captcha.NewCaptcha().GenerateCaptcha()

	logger.LogIf(err)

	type resp struct {
		captchaId     string `json:"captcha_id"`
		captchaBase64 string `json:"captcha_base64"`
	}

	//c.JSON(http.StatusOK, gin.H{
	//	"captcha_id":    id,
	//	"captcha_image": base64,
	//})

	response.Success(c, gin.H{
		"captcha_id":    id,
		"captcha_image": base64,
	})
}

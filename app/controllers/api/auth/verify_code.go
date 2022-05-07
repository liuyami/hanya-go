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

	response.Success(c, gin.H{
		"captcha_id":    id,
		"captcha_image": base64,
	})
}

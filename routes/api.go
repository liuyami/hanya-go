package routes

import (
	"hanya-go/app/controllers/api/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {

	root := r.Group("/api")
	{
		root.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		})

		userAuthGroup := root.Group("/auth")
		{
			//suc := new(auth.Signup)

			userAuthGroup.POST("/signup/phone/exist", auth.IsPhoneExist)
			userAuthGroup.POST("/signup/email/exist", auth.IsEmailExist)

			// 图片验证码
			userAuthGroup.GET("/verify-codes/captcha", auth.GetCaptcha)
		}
	}

}

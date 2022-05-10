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
			userAuthGroup.POST("/signup/using-phone", auth.SignupUsingPhone)
			userAuthGroup.POST("/signup/using-email", auth.SignupUsingEmail)

			// 图片验证码
			userAuthGroup.GET("/verify-codes/captcha", auth.GetCaptcha)
			// 发送短信验证码
			userAuthGroup.POST("/verify-codes/phone", auth.SendUsingPhone)
			userAuthGroup.POST("/verify-codes/email", auth.SendUsingEmail)

			// 扽牢固
			userAuthGroup.POST("/login/using-phone", auth.LoginByPhone)
			userAuthGroup.POST("/login/using-password", auth.LoginByPassword)
			userAuthGroup.POST("/login/refresh-token", auth.RefreshToken)

			// 重置密码
			userAuthGroup.POST("/password-reset/using-phone", auth.ResetPasswordByPhone)
			userAuthGroup.POST("/password-reset/using-email", auth.ResetPasswordByEmail)
		}
	}

}

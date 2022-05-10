package routes

import (
	"hanya-go/app/controllers/api/auth"
	"hanya-go/app/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(r *gin.Engine) {

	apiGroup := r.Group("/api")
	apiGroup.Use(middlewares.LimitIP("200-H"))
	{
		apiGroup.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		})

		userAuthGroup := apiGroup.Group("/auth")
		{
			//suc := new(auth.Signup)
			//检测是否存在
			userAuthGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), auth.IsPhoneExist)
			userAuthGroup.POST("/signup/email/exist", middlewares.GuestJWT(), auth.IsEmailExist)
			//注册
			userAuthGroup.POST("/signup/using-phone", middlewares.GuestJWT(), auth.SignupUsingPhone)
			userAuthGroup.POST("/signup/using-email", middlewares.GuestJWT(), auth.SignupUsingEmail)

			// 图片验证码
			userAuthGroup.GET("/verify-codes/captcha", middlewares.GuestJWT(), auth.GetCaptcha)
			// 发送短信验证码
			userAuthGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-H"), auth.SendUsingPhone)
			userAuthGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), auth.SendUsingEmail)

			// 登录
			userAuthGroup.POST("/login/using-phone", middlewares.GuestJWT(), auth.LoginByPhone)
			userAuthGroup.POST("/login/using-password", middlewares.GuestJWT(), auth.LoginByPassword)
			userAuthGroup.POST("/login/refresh-token", middlewares.AuthJWT(), auth.RefreshToken)

			// 重置密码
			userAuthGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), auth.ResetPasswordByPhone)
			userAuthGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), auth.ResetPasswordByEmail)
		}
	}

}

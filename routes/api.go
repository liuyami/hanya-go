package routes

import (
	"github.com/gin-gonic/gin"
	"hanya-go/app/controllers/api/auth"
	"hanya-go/app/controllers/api/category"
	"hanya-go/app/controllers/api/link"
	"hanya-go/app/controllers/api/topic"
	"hanya-go/app/controllers/api/user"
	"hanya-go/app/middlewares"
	"net/http"
)

func RegisterAPIRoutes(r *gin.Engine) {

	apiGroup := r.Group("/api")

	apiGroup.GET("/user", middlewares.AuthJWT(), user.CurrentUser)

	apiGroup.Use(middlewares.LimitIP("200-H"))
	{
		apiGroup.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		})

		// 用户
		userGroup := apiGroup.Group("/users")
		{
			userGroup.GET("", user.Index)
			userGroup.PUT("", middlewares.AuthJWT(), user.UpdateProfile)
			userGroup.POST("/avatar", middlewares.AuthJWT(), user.UpdateAvatar)
		}

		// 分类
		categoryGroup := apiGroup.Group("/categories")
		{
			categoryGroup.GET("", category.Index)
			categoryGroup.POST("", middlewares.AuthJWT(), category.Store)
			categoryGroup.PUT("/:category_id", middlewares.AuthJWT(), category.Update)
			categoryGroup.DELETE("/:category_id", middlewares.AuthJWT(), category.Delete)
		}

		// 话题
		topicGroup := apiGroup.Group("/topic")
		{
			topicGroup.GET("", topic.Index)
			topicGroup.GET("/:topic_id", topic.Show)
			topicGroup.POST("", middlewares.AuthJWT(), topic.Store)
			topicGroup.PUT("/:topic_id", middlewares.AuthJWT(), topic.Update)
			topicGroup.POST("/:topic_id", middlewares.AuthJWT(), topic.Delete)
		}

		// 友情链接
		linkGroup := apiGroup.Group("/link")
		{
			linkGroup.GET("", link.Index)
		}

		// 账号相关
		authGroup := apiGroup.Group("/auth")
		{
			//suc := new(auth.Signup)
			//检测是否存在
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), auth.IsPhoneExist)
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), auth.IsEmailExist)
			//注册
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), auth.SignupUsingPhone)
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), auth.SignupUsingEmail)

			// 图片验证码
			authGroup.GET("/verify-codes/captcha", middlewares.GuestJWT(), auth.GetCaptcha)
			// 发送短信验证码
			authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-H"), auth.SendUsingPhone)
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("20-H"), auth.SendUsingEmail)

			// 登录
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), auth.LoginByPhone)
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), auth.LoginByPassword)
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), auth.RefreshToken)

			// 重置密码
			authGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), auth.ResetPasswordByPhone)
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), auth.ResetPasswordByEmail)
		}

	}

}

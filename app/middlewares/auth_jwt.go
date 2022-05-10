package middlewares

import (
	"github.com/gin-gonic/gin"
	"hanya-go/app/models/user"
	"hanya-go/app/response"
	"hanya-go/pkg/jwt"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从标头 Authorization:Bearer xxxxx 中获取信息，并验证 JWT 的准确性
		claims, err := jwt.NewJWT().ParserToken(c)
		if err != nil {
			response.App(c, 401, "认证失败，请登录")
			return
		}

		// jwt 解析成功，设置用户信息
		userModel := user.GetById(claims.UserID)
		if userModel.UserID == 0 {
			response.App(c, 401, "账号异常，用户可能被警用或删除")
		}

		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
		c.Set("current_user_id", userModel.GetStringID())
		c.Set("current_user", userModel)

		c.Next()
	}
}

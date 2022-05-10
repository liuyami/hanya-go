package middlewares

import (
	"github.com/gin-gonic/gin"
	"hanya-go/app/response"
	"hanya-go/pkg/jwt"
)

func GuestJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.GetHeader("authorization")) > 0 {
			// 解析 token 成功，说明登录成功了
			_, err := jwt.NewJWT().ParserToken(c)
			if err == nil {
				response.App(c, 401, "请先退出登录以游客身份访问")
				c.Abort()
				return
			}
		}

		c.Next()
	}
}

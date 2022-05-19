package middlewares

import (
	"github.com/gin-gonic/gin"
	"hanya-go/app/response"
)

// ForceUA 中间件，强制请求必须附带 User-Agent 标头
func ForceUA() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取 header 中的 user-agent 标头信息

		if len(c.Request.Header["User-Agent"]) == 0 {
			response.App(c, 400, "请求必须附带 User-Agent 标头")
			return
		}

		c.Next()
	}
}

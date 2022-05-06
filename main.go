package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	r.NoRoute(func(ctx *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := ctx.Request.Header.Get("Accept")

		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			ctx.String(http.StatusNotFound, "Page not found")
		} else {
			// 默认返回 JSON
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}

	})

	r.Run(":8000")
}

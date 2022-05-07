package bootstrap

import (
	"hanya-go/app/middlewares"
	"net/http"
	"strings"

	"hanya-go/routes"

	"github.com/gin-gonic/gin"
)

func SetRoute(router *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleWare(router)

	// 注册API路由
	routes.RegisterAPIRoutes(router)

	// 注册 404 路由
	setup404Handle(router)

}

// 注册全局中间件
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(
		middlewares.Logger(),
		gin.Recovery(),
	)
}

func setup404Handle(router *gin.Engine) {
	router.NoRoute(func(ctx *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := ctx.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			ctx.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			ctx.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}

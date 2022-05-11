package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"hanya-go/bootstrap"
	"hanya-go/pkg/config"
	"hanya-go/pkg/console"
	"hanya-go/pkg/logger"
)

var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "start web server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {
	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	// gin 实例
	router := gin.New()

	// 初始化路由
	bootstrap.SetRoute(router)

	// 运行服务
	err := router.Run(":" + config.Get("app.port"))

	if err != nil {
		logger.ErrorString("cmd", "serve", err.Error())
		console.Exit("unable to start serve, error:" + err.Error())
	}

}

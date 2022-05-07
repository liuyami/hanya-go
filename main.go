package main

import (
	"flag"
	"fmt"

	"hanya-go/bootstrap"
	btsConfig "hanya-go/config"
	"hanya-go/pkg/config"

	"github.com/gin-gonic/gin"
)

func init() {
	btsConfig.Initialize()
}

func main() {
	// 配置初始化，依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件， 如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// 初始化日志
	bootstrap.SetupLogger()

	router := gin.New()

	// 初始化数据库
	bootstrap.SetupDB()

	// 初始化路由
	bootstrap.SetRoute(router)

	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}

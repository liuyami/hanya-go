package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"hanya-go/bootstrap"
	btsConfig "hanya-go/config"
	"hanya-go/pkg/config"
	"hanya-go/pkg/verifycode"
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

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	// 初始化数据库
	bootstrap.SetupDB()

	// redis
	bootstrap.SetupRedis()

	// 初始化路由
	bootstrap.SetRoute(router)

	//logger.Dump(captcha.NewCaptcha().VerifyCaptcha("hCyFVozJt7Djt6vStbYz", "6291"), "正确的答案")
	//logger.Dump(captcha.NewCaptcha().VerifyCaptcha("hCyFVozJt7Djt6vStbYz", "1234"), "错误的答案")

	//sms.NewSMS().Send("18107397886", sms.Message{
	//	Template: config.GetString("sms.aliyun.template_code"),
	//	Data:     map[string]string{"code": "1234"},
	//})

	verifycode.NewVerifyCode().SendSMS("18107397886")

	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}

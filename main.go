package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"hanya-go/app/cmd"
	"hanya-go/app/cmd/make"
	"hanya-go/bootstrap"
	btsConfig "hanya-go/config"
	"hanya-go/pkg/config"
	"hanya-go/pkg/console"
	"os"
)

func init() {
	btsConfig.Initialize()
}

func main() {

	// 应用的入口，默认调用 cmd.serve

	var rootCmd = &cobra.Command{
		Use:   config.Get("app.name"),
		Short: "Yami test golang project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {
			// 配置初始化，依赖命令行 --env 参数
			config.InitConfig(cmd.Env)

			// 初始化 Logger
			bootstrap.SetupLogger()

			// 初始化数据库
			bootstrap.SetupDB()

			// 初始化 Redis
			bootstrap.SetupRedis()

			// 初始化缓存
			bootstrap.SetupCache()

		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
		make.CmdMake,
		cmd.CmdMigrate,
		cmd.CmdDBSeed,
		cmd.CmdCache,
	)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}

}

package make

import (
	"fmt"
	"github.com/spf13/cobra"
)

var CmdMakeAPI = &cobra.Command{
	Use:   "api",
	Short: "Create api controller，exmaple: make api auth/user",
	Run:   runMakeAPI,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeAPI(cmd *cobra.Command, args []string) {
	model := makeModelFromString(args[0])

	// 组建目标目录
	filePath := fmt.Sprintf("app/controllers/api/%s.go", model.TableName)

	// 基于模板创建文件（做好变量替换）
	createFileFromStub(filePath, "api", model)
}

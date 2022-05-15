package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var CmdMakeAPI = &cobra.Command{
	Use:   "api",
	Short: "Create api controller，exmaple: make api auth/user",
	Run:   runMakeAPI,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakeAPI(cmd *cobra.Command, args []string) {
	savePath := "app/controllers/api"
	fileName := ""

	if strings.Contains(args[0], "/") {
		pathList := strings.Split(args[0], "/")
		pathListLen := len(pathList)

		for i := 0; i < pathListLen-1; i++ {
			savePath = savePath + "/" + pathList[i]
		}

		_ = os.MkdirAll(savePath, os.ModePerm)
		fileName = pathList[pathListLen-1]
	}

	model := makeModelFromString(fileName)

	// 组建目标目录
	filePath := fmt.Sprintf("%s/%s.go", savePath, model.TableName)

	// 基于模板创建文件（做好变量替换）
	createFileFromStub(filePath, "api", model)
}

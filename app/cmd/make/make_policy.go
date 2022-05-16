package make

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var CmdMakePolicy = &cobra.Command{
	Use:   "policy",
	Short: "Create policy file, example: make policy user",
	Run:   runMakePolicy,
	Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runMakePolicy(cmd *cobra.Command, args []string) {

	savePath := "app/policy"
	fileName := args[0]

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

	fmt.Println(model)

	// 拼接目标文件路径
	filePath := fmt.Sprintf("app/policy/%s_policy.go", model.PackageName)

	// 基于模板创建文件（做好变量替换）
	createFileFromStub(filePath, "policy", model)
}

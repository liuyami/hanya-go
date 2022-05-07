package requests

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"hanya-go/app/response"
	"net/http"
)

// ValidatorFunc 验证函数类型
type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

func Validate(c *gin.Context, obj interface{}, handle ValidatorFunc) bool {
	// 1. 解析请求，支持 JSON 数据、表单请求和 URL Query
	if err := c.ShouldBind(obj); err != nil {

		response.App(c, http.StatusBadRequest, "请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")

		fmt.Println(err.Error())

		return false
	}

	// 2. 验证表单
	errs := handle(obj, c)

	// 3. 判断验证是否通过
	if len(errs) > 0 {
		response.Fail(c, http.StatusUnprocessableEntity, "请求参数错误", errs)
		return false
	}

	return true
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	// 配置选项
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid", // 模型中的 Struct 标签标识符
		Messages:      messages,
	}

	// 开始验证
	return govalidator.New(opts).ValidateStruct()
}
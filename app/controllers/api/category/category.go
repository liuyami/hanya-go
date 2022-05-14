package category

import (
	"hanya-go/app/models/category"
	"hanya-go/app/requests"
	"hanya-go/app/response"

	"github.com/gin-gonic/gin"
)

func Store(c *gin.Context) {

	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}

	categoryModel := category.Category{
		Name: request.Name,
		Desc: request.Desc,
	}

	categoryModel.Create()

	if categoryModel.CategoryID > 0 {
		response.Success(c, categoryModel)
	} else {
		response.Fail(c, 1001, "创建失败，请稍后尝试~", nil)
	}
}

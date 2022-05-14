package category

import (
	"hanya-go/app/models/category"
	"hanya-go/app/requests"
	"hanya-go/app/response"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	request := requests.CategoryIndexRequest{}
	if ok := requests.Validate(c, &request, requests.CategoryIndexFun); !ok {
		return
	}

	sort := c.DefaultQuery("sort", "category_id")
	order := c.DefaultQuery("order", "desc")

	data, pager := category.Paginate(c, 3, sort, order, "/api/categories")

	response.Success(c, gin.H{
		"list":  data,
		"pager": pager,
	})
}

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

func Update(c *gin.Context) {

	// 验证 url 参数 id 是否正确
	categoryModel := category.Get(c.Param("category_id"))
	if categoryModel.CategoryID == 0 {
		response.Fail(c, 1001, "分类不存在", nil)
		return
	}

	// 表单验证
	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}

	// 保存数据
	categoryModel.Name = request.Name
	categoryModel.Desc = request.Desc
	rowsAffected := categoryModel.Save()

	if rowsAffected > 0 {
		response.Success(c, categoryModel)
	} else {
		response.Fail(c, 1002, "更新失败", nil)
	}
}

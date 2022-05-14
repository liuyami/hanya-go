package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type CategoryRequest struct {
	Name string `valid:"name" json:"name"`
	Desc string `valid:"desc" json:"desc,omitempty"`
}

func CategorySave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name":        []string{"required", "min_cn:2", "max_cn:8", "not_exists:categories,name"},
		"description": []string{"min_cn:3", "max_cn:255"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:分类名称为必填项",
			"min_cn:分类名称长度需至少 2 个字",
			"max_cn:分类名称长度不能超过 8 个字",
			"not_exists:分类名称已存在",
		},
		"description": []string{
			"min_cn:分类描述长度需至少 3 个字",
			"max_cn:分类描述长度不能超过 255 个字",
		},
	}
	return validate(data, rules, messages)
}

// 分页

type CategoryIndexRequest struct {
	Sort    string `valid:"sort" form:"sort"`
	Order   string `valid:"order" form:"order"`
	PerPage string `valid:"per_page" form:"per_page"`
}

func CategoryIndexFun(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"sort":     []string{"in:category_id,created_at,updated_at"},
		"order":    []string{"in:asc,desc"},
		"per_page": []string{"numeric_between:2,100"},
	}

	messages := govalidator.MapData{
		"sort": []string{
			"in:排序字段仅支持 category_id,created_at,updated_at",
		},
		"order": []string{
			"in:排序规则仅支持 asc（正序）,desc（倒序）",
		},
		"per_page": []string{
			"numeric_between:每页条数的值介于 2~100 之间",
		},
	}

	return validate(data, rules, messages)
}

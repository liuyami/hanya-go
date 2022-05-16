package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type TopicRequest struct {
	Title      string `json:"title,omitempty" valid:"title"`
	Body       string `json:"body,omitempty" valid:"body"`
	CategoryID uint64 `json:"category_id,omitempty" valid:"category_id"`
}

func TopicSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"title":       []string{"required", "min_cn:3", "max_cn:40"},
		"body":        []string{"required", "min_cn:10", "max_cn:50000"},
		"category_id": []string{"required", "exists:category,category_id"},
	}
	
	messages := govalidator.MapData{
		"title": []string{
			"required:帖子标题为必填项",
			"min_cn:标题长度需大于 3",
			"max_cn:标题长度需小于 40",
		},
		"body": []string{
			"required:帖子内容为必填项",
			"min_cn:长度需大于 10",
		},
		"category_id": []string{
			"required:帖子分类为必填项",
			"exists:帖子分类未找到",
		},
	}
	return validate(data, rules, messages)
}

// TopicIndexRequest 分页
type TopicIndexRequest struct {
	Sort    string `valid:"sort" form:"sort"`
	Order   string `valid:"order" form:"order"`
	PerPage string `valid:"per_page" form:"per_page"`
}

func TopicIndexFun(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"sort":     []string{"in:topic_id,created_at,updated_at"},
		"order":    []string{"in:asc,desc"},
		"per_page": []string{"numeric_between:2,100"},
	}

	messages := govalidator.MapData{
		"sort": []string{
			"in:排序字段仅支持 topic_id,created_at,updated_at",
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

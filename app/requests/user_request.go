package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type UserIndexRequest struct {
	Sort    string `valid:"sort" form:"sort"`
	Order   string `valid:"order" form:"order"`
	PerPage string `valid:"per_page" form:"per_page"`
}

func UserIndexFun(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"sort":     []string{"in:user_id,created_at,updated_at"},
		"order":    []string{"in:asc,desc"},
		"per_page": []string{"numeric_between:2,100"},
	}

	messages := govalidator.MapData{
		"sort": []string{
			"in:排序字段仅支持 id,created_at,updated_at",
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

type UserUpdateProfileRequest struct {
	Nickname string `valid:"nickname" json:"nickname"`
	Avatar   string `valid:"avatar" json:"avatar"`
}

func UserUpdateProfileRequestFun(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"nickname": []string{"required"},
		"avatar":   []string{"required"},
	}

	messages := govalidator.MapData{
		"nickname": []string{
			"required:昵称必须填写",
		},
		"avatar": []string{
			"required:必须有头像",
		},
	}

	return validate(data, rules, messages)
}

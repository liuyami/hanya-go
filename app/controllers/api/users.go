package api

import (
	"hanya-go/app/models/user"
	"hanya-go/app/requests"
	"hanya-go/app/response"
	"hanya-go/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Success(c, userModel)
}

func Index(c *gin.Context) {

	request := requests.UserIndexRequest{}

	if ok := requests.Validate(c, &request, requests.UserIndexFun); !ok {
		return
	}

	sort := c.DefaultQuery("sort", "user_id")
	order := c.DefaultQuery("order", "desc")

	data, pager := user.Paginate(c, 3, sort, order)

	response.Success(c, gin.H{
		"list":  data,
		"pager": pager,
	})
}

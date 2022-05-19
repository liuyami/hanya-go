package user

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

func UpdateProfile(c *gin.Context) {
	req := requests.UserUpdateProfileRequest{}

	if ok := requests.Validate(c, &req, requests.UserUpdateProfileRequestFun); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Nickname = req.Nickname
	currentUser.Avatar = req.Avatar
	rowsAffected := currentUser.Save()

	if rowsAffected > 0 {
		response.Success(c, nil)
	} else {
		response.Fail(c, 1001, "更新失败，请稍后再试", nil)
	}
}

package user

import (
	"hanya-go/app/models/user"
	"hanya-go/app/requests"
	"hanya-go/app/response"
	"hanya-go/pkg/auth"
	"hanya-go/pkg/config"
	"hanya-go/pkg/file"

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
	rowsAffected := currentUser.Save()

	if rowsAffected > 0 {
		response.Success(c, nil)
	} else {
		response.Fail(c, 1001, "更新失败，请稍后再试", nil)
	}
}

func UpdateAvatar(c *gin.Context) {
	request := requests.UserUpdateAvatarRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateAvatarHandle); !ok {
		return
	}

	// 保存图片
	avatar, err := file.SaveUploadAvatar(c, request.Avatar)
	if err != nil {
		response.Fail(c, 1001, "上传失败，请稍后再试试", nil)
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Avatar = config.GetString("app.url") + avatar
	currentUser.Save()

	response.Success(c, currentUser)
}

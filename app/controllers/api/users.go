package api

import (
	"hanya-go/app/models/user"
	"hanya-go/app/response"
	"hanya-go/pkg/auth"

	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Success(c, userModel)
}

func Index(c *gin.Context) {
	data := user.All()

	response.Success(c, data)
}

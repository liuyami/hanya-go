package link

import (
	"hanya-go/app/models/link"
	"hanya-go/app/response"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	link := link.AllCached()
	response.Success(c, link)
}

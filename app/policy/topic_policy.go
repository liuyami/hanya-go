package policy

import (
	"github.com/gin-gonic/gin"
	"hanya-go/app/models/topic"
	"hanya-go/pkg/auth"
)

func CanModifyTopic(c *gin.Context, _topic topic.Topic) bool {
	return auth.CurrentUserId(c) == _topic.UserID
}

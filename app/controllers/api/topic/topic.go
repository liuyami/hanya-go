package topic

import (
	"github.com/gin-gonic/gin"
	"hanya-go/app/models/topic"
	"hanya-go/app/requests"
	"hanya-go/app/response"
	"hanya-go/pkg/auth"
)

func Index(c *gin.Context) {
	request := requests.TopicIndexRequest{}
	if ok := requests.Validate(c, &request, requests.TopicIndexFun); !ok {
		return
	}

	sort := c.DefaultQuery("sort", "topic_id")
	order := c.DefaultQuery("order", "desc")

	data, pager := topic.Paginate(c, 3, sort, order)

	response.Success(c, gin.H{
		"list":  data,
		"pager": pager,
	})
}

func Store(c *gin.Context) {

	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	topicModel := topic.Topic{
		Title:      request.Title,
		Body:       request.Body,
		CategoryID: request.CategoryID,
		UserID:     auth.CurrentUserId(c),
	}

	topicModel.Create()

	if topicModel.TopicID > 0 {
		response.Success(c, nil)
	} else {
		response.Fail(c, 1001, "创建失败，请稍后尝试~", nil)
	}
}

package topic

import (
	"github.com/gin-gonic/gin"
	"hanya-go/app/models/topic"
	"hanya-go/app/policy"
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

func Show(c *gin.Context) {
	topicModel := topic.Get(c.Param("topic_id"))
	if topicModel.TopicID == 0 {
		response.App(c, 404, "记录不存在")
		return
	}
	response.Success(c, topicModel)
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

func Update(c *gin.Context) {
	topicModel := topic.Get(c.Param("topic_id"))
	if topicModel.TopicID == 0 {
		response.Fail(c, 1001, "记录没找到", nil)
		return
	}

	if ok := policy.TopicAuthCheck(c, topicModel); !ok {
		response.App(c, 403, "没有权限")
		return
	}

	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	topicModel.Title = request.Title
	topicModel.Body = request.Body
	topicModel.CategoryID = request.CategoryID
	rowsAffected := topicModel.Save()
	if rowsAffected > 0 {
		response.Success(c, topicModel)
	} else {
		response.Fail(c, 1002, "更新失败，请稍后尝试~", nil)
	}
}

func Delete(c *gin.Context) {
	topicModel := topic.Get(c.Param("topic_id"))
	if topicModel.TopicID == 0 {
		response.Fail(c, 1001, "记录没找到", nil)
		return
	}

	if ok := policy.TopicAuthCheck(c, topicModel); !ok {
		response.App(c, 403, "没有权限")
		return
	}

	rowsAffected := topicModel.Delete()

	if rowsAffected > 0 {
		response.Success(c, nil)
		return
	}

	response.Fail(c, 1001, "删除失败，请稍后再试", nil)
}

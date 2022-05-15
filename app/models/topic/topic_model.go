// Package topic 模型
package topic

import (
	"hanya-go/app/models"
	"hanya-go/app/models/category"
	"hanya-go/app/models/user"
	"hanya-go/pkg/app"
	"hanya-go/pkg/database"
	"hanya-go/pkg/paginator"

	"github.com/gin-gonic/gin"
)

type Topic struct {
	TopicID uint `gorm:"column:topic_id;primaryKey;autoIncrement;type:int(10)" json:"topic_id,omitempty"`

	Title      string `json:"title,omitempty" `
	Body       string `json:"body,omitempty" `
	UserID     uint   `json:"user_id,omitempty"`
	CategoryID uint   `json:"category_id,omitempty"`

	// 通过 user_id 关联用户
	User user.User `json:"user"`

	// 通过 category_id 关联分类
	Category category.Category `json:"category"`

	models.Datetime
}

func (Topic) TableName() string {
	return "topic"
}

func (topic *Topic) Create() {
	database.DB.Create(&topic)
}

func (topic *Topic) Save() (rowsAffected int64) {
	result := database.DB.Save(&topic)
	return result.RowsAffected
}

func (topic *Topic) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&topic)
	return result.RowsAffected
}

// Paginate 页面获取
func Paginate(c *gin.Context, perPage int, sort, order string) (topic []Topic, paging paginator.Paging) {

	paging = paginator.Paginate(
		c,
		database.DB.Model(Topic{}),
		&topic,
		app.URL("/api/topic"),
		perPage,
		sort,
		order,
	)

	return
}

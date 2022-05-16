// Package link 模型
package link

import (
	"hanya-go/app/models"
	"hanya-go/pkg/app"
	"hanya-go/pkg/database"
	"hanya-go/pkg/paginator"

	"github.com/gin-gonic/gin"
)

type Link struct {
	LinkID uint64 `gorm:"column:link_id;primaryKey;autoIncrement;type:int(10)" json:"link_id,omitempty"`

	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`

	models.Datetime
}

func (Link) TableName() string {
	return "link"
}

func (link *Link) Create() {
	database.DB.Create(&link)
}

func (link *Link) Save() (rowsAffected int64) {
	result := database.DB.Save(&link)
	return result.RowsAffected
}

func (link *Link) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&link)
	return result.RowsAffected
}

// Paginate 页面获取
func Paginate(c *gin.Context, perPage int, sort, order string) (link []Link, paging paginator.Paging) {

	paging = paginator.Paginate(
		c,
		database.DB.Model(Link{}),
		&link,
		app.URL("/api/link"),
		perPage,
		sort,
		order,
	)

	return
}

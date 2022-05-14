// Package category 模型
package category

import (
	"hanya-go/app/models"
	"hanya-go/pkg/app"
	"hanya-go/pkg/database"
	"hanya-go/pkg/paginator"

	"github.com/gin-gonic/gin"
)

type Category struct {
	CategoryID uint `gorm:"column:category_id;primaryKey;autoIncrement;type:int(10)" json:"category_id,omitempty"`

	Name string `json:"name,omitempty"`
	Desc string `json:"desc,omitempty"`

	models.Datetime
}

func (Category) TableName() string {
	return "category"
}

func (category *Category) Create() {
	database.DB.Create(&category)
}

func (category *Category) Save() (rowsAffected int64) {
	result := database.DB.Save(&category)
	return result.RowsAffected
}

func (category *Category) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&category)
	return result.RowsAffected
}

// Paginate 页面获取
func Paginate(c *gin.Context, perPage int, sort, order string) (category []Category, paging paginator.Paging) {

	paging = paginator.Paginate(
		c,
		database.DB.Model(Category{}),
		&category,
		app.URL("/api/category"),
		perPage,
		sort,
		order,
	)

	return
}

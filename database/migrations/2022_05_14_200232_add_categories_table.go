package migrations

import (
	"database/sql"
	"hanya-go/app/models"
	"hanya-go/app/models/category"
	"hanya-go/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		CategoryID uint `gorm:"column:category_id;primaryKey;autoIncrement;type:int(10)" json:"category_id,omitempty"`

		Name string `json:"name,omitempty"`
		Desc string `json:"desc,omitempty"`

		models.Datetime
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		_ = migrator.AutoMigrate(&category.Category{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		_ = migrator.DropTable(&category.Category{})
	}

	migrate.Add("2022_05_14_200232_add_categories_table", up, down)
}

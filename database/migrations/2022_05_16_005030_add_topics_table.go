package migrations

import (
	"database/sql"
	"hanya-go/app/models"
	"hanya-go/app/models/topic"
	"hanya-go/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		UserID uint `gorm:"column:user_id;primaryKey;autoIncrement;type:int(10)"`
	}
	type Category struct {
		CategoryID uint `gorm:"column:category_id;primaryKey;autoIncrement;type:int(10)"`
	}

	type Topic struct {
		TopicID uint `gorm:"column:topic_id;primaryKey;autoIncrement;type:int(10)"`

		Title      string `gorm:"type:varchar(255);not null;index"`
		Body       string `gorm:"type:text;not null"`
		UserID     uint   `gorm:"column:user_id;type:int(10);not null;index"`
		CategoryID uint   `gorm:"column:category_id;type:int(10);not null;index"`

		// 会创建 user_id 和 category_id 外键的约束
		User     User
		Category Category

		models.Datetime
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		_ = migrator.AutoMigrate(&topic.Topic{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		_ = migrator.DropTable(&topic.Topic{})
	}

	migrate.Add("2022_05_16_005030_add_topics_table", up, down)
}

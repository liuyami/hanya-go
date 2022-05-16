package migrations

import (
	"database/sql"
	"hanya-go/app/models"
	"hanya-go/app/models/link"
	"hanya-go/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		LinkID uint `gorm:"column:link_id;primaryKey;autoIncrement;type:int(10)" json:"link_id,omitempty"`

		Name string `gorm:"type:varchar(255);not null"`
		URL  string `gorm:"type:varchar(255);default:null"`
		models.Datetime
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		_ = migrator.AutoMigrate(&link.Link{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		_ = migrator.DropTable(&link.Link{})
	}

	migrate.Add("2022_05_17_012528_add_links_table", up, down)
}

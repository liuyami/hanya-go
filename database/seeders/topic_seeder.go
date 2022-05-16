package seeders

import (
	"fmt"
	"hanya-go/database/factories"
	"hanya-go/pkg/console"
	"hanya-go/pkg/logger"
	"hanya-go/pkg/seed"

	"gorm.io/gorm"
)

func init() {

	seed.Add("SeedTopicTable", func(db *gorm.DB) {

		topic := factories.MakeTopic(50)

		result := db.Table("topic").Create(&topic)

		if err := result.Error; err != nil {
			logger.LogIf(err)
			return
		}

		console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
	})
}

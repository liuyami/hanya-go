package topic

import (
	"hanya-go/pkg/database"
)

func Get(idStr string) (topic Topic) {
	database.DB.Where("topic_id", idStr).First(&topic)
	return
}

func GetBy(field, value string) (topic Topic) {
	database.DB.Where("? = ?", field, value).First(&topic)
	return
}

func All() (topic []Topic) {
	database.DB.Find(&topic)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Topic{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

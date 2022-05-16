package factories

import (
	"github.com/bxcodec/faker/v3"
	"hanya-go/app/models/topic"
)

func MakeTopic(count int) []topic.Topic {

	var objs []topic.Topic

	// 设置唯一性，如 Topic 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		topicModel := topic.Topic{
			Title:      faker.Sentence(),
			Body:       faker.Paragraph(),
			CategoryID: 3,
			UserID:     1,
		}
		objs = append(objs, topicModel)
	}

	return objs
}

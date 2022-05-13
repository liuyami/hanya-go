package factories

import (
	"github.com/bxcodec/faker/v3"
	"hanya-go/app/models/user"
	"hanya-go/pkg/helpers"
)

func MakeUsers(times int) []user.User {
	var userList []user.User

	// 设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i := 0; i < times; i++ {
		model := user.User{
			Nickname: faker.Name(),
			Avatar:   "",
			Openid:   helpers.RandomString(32),
			Email:    faker.Email(),
			Phone:    helpers.RandomNumber(11),
			Password: "$2a$14$oPzVkIdwJ8KqY0erYAYQxOuAAlbI/sFIsH0C0R4MPc.3JbWWSuaUe",
		}

		userList = append(userList, model)
	}

	return userList
}

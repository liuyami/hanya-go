package user

import (
	"gorm.io/gorm"
	"hanya-go/pkg/hash"
)

// BeforeSave 在模型创建和更新前被调用，对密码做加密
func (userModel *User) BeforeSave(tx *gorm.DB) (err error) {
	if !hash.BcryptIsHashed(userModel.Password) {
		userModel.Password = hash.BcryptHash(userModel.Password)
	}

	return
}

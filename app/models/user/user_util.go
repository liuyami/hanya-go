package user

import "hanya-go/pkg/database"

// IsEmailExist 判断EMAIL是否存在
func IsEmailExist(email string) bool {
	var count int64

	database.DB.Model(&User{}).Where("email = ?", email).Count(&count)

	return count > 0
}

// IsPhoneExist 判断手机号码是否存在
func IsPhoneExist(phone string) bool {
	var count int64

	database.DB.Model(&User{}).Where("phone = ?", phone).Count(&count)

	return count > 0
}

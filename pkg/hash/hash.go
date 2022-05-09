package hash

import (
	"golang.org/x/crypto/bcrypt"
	"hanya-go/pkg/logger"
)

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(str string) string {
	// GenerateFromPassword 的第二个参数是 cost 值。建议大于 12，数值越大耗费时间越长
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), 14)
	logger.LogIf(err)

	return string(bytes)
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))

	return err == nil
}

// BcryptIsHashed 判断字符串是否是哈希过的数据
func BcryptIsHashed(str string) bool {
	// bcrypt 加密后的长度等于 60
	return len(str) == 60
}

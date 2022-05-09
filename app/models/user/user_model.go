package user

import (
	"github.com/spf13/cast"
	"hanya-go/app/models"
	"hanya-go/pkg/database"
	"hanya-go/pkg/hash"
)

type User struct {
	// models.ID

	UserID uint `gorm:"column:user_id;primaryKey;autoIncrement;type:int(10)" json:"user_id,omitempty"`

	Nickname string `gorm:"column:nickname;type:varchar(128)" json:"nickname,omitempty"`
	Avatar   string `gorm:"column:avatar;type:varchar(256)" json:"avatar,omitempty"`
	Openid   string `gorm:"column:openid;uniqueIndex;type:varchar(128)" json:"openid,omitempty"`
	Email    string `gorm:"uniqueIndex;type:varchar(128)" json:"-"`
	Phone    string `gorm:"uniqueIndex;type:char(11)" json:"-"`
	Password string `gorm:"type:varchar(256)" json:"-"`

	models.Datetime
}

func (User) TableName() string {
	return "users"
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(userModel.Password, _password)
}

// GetStringID 获取 User ID 的字符串格式
func (userModel *User) GetStringID() string {
	return cast.ToString(userModel.UserID)
}

package user

import "hanya-go/app/models"

type User struct {
	models.ID

	NickName string `gorm:"column:nickname;type:varchar(128)" json:"nickname,omitempty"`
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

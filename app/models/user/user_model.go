package user

import "hanya-go/app/models"

type User struct {
	models.ID

	Name     string `json:"name,omitempty"`
	Email    string `gorm:"uniqueIndex" json:"-"`
	Phone    string `gorm:"uniqueIndex" json:"-"`
	Password string `json:"-"`

	models.Datetime
}

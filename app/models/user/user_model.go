package user

import "hanya-go/app/models"

type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"uniqueIndex,-"`
	Phone    string `json:"uniqueIndex,-"`
	Password string `json:"-"`

	models.CommonTimestampsField
}

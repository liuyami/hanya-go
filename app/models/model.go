package models

import "time"

type ID struct {
	ID uint `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

type Datetime struct {
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;" json:"updated_at,omitempty"`
}

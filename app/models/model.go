package models

import "time"

type ID struct {
	ID uint `gorm:"column:id;primaryKey;autoIncrement;type:int(10)" json:"id,omitempty"`
}

type Datetime struct {
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;type:datetime(0)" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime;type:datetime(0)" json:"updated_at,omitempty"`
}

package model

import "time"

type Tag struct {
	Id        int       `json:"id" gorm:"column:id;"`
	Name      string    `json:"name" gorm:"column:name;"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Tag) TableName() string {
	return "Tags"
}

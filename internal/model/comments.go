package model

import "time"

type Comments struct {
	Id        int       `json:"id" gorm:"column:id"`
	UserId    int       `json:"user_id" gorm:"user_id"`
	Content   string    `json:"content" gorm:"content"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Comments) TableName() string {
	return "comments"
}

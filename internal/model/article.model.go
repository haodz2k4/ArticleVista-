package model

import "time"

type Articles struct {
	Id        int       `json:"id" gorm:"column:id;"`
	Title     string    `json:"title" gorm:"column:title;"`
	Content   string    `json:"content" gorm:"column:content;"`
	UserId    int       `json:"user_id" gorm:"column:user_id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Articles) TableName() string {
	return "Articles"
}

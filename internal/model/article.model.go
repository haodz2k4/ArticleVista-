package model

import "time"

type Articles struct {
	Id        int       `json:"id" gorm:"column:id;"`
	Title     string    `json:"title" gorm:"column:title;" validate:"required,min=2"`
	Content   string    `json:"content" gorm:"column:content;" validate:"required,min=5"`
	UserId    int       `json:"user_id" gorm:"column:user_id" validate:"required,gt=1"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Articles) TableName() string {
	return "Articles"
}

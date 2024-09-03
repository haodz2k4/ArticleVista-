package model

import "time"

type Users struct {
	Id        int       `json:"id" gorm:"column:id;"`
	UserName  string    `json:"user_name" gorm:"column:userName"`
	Email     string    `json:"email" gorm:"column:email"`
	Password  string    `json:"password" gorm:"column:password"`
	BirthDate time.Time `json:"birth_date" gorm:"column:birthDate"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at"`
}

func (Users) TableName() string {
	return "Users"
}

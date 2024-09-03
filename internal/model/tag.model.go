package model

type Tag struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
}

func (Tag) TableName() string {
	return "Tags"
}

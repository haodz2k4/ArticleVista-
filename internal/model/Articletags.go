package model

type Articletags struct {
	ArticleId int `json:"article_id" gorm:"column:article_id"`
	TagId     int `json:"tag_id" gorm:"column:tag_id"`
}

func (Articletags) TableName() string {
	return "Articletags"
}

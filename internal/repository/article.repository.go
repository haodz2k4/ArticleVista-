package repository

import (
	"ArticleVista/internal/model"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	CreateArticle(article *model.Articles) error
	GetArticleById(id int) (*model.Articles, error)
	GetAllArticles() ([]model.Articles, error)
	UpdateArticle(article *model.Articles) error
	DeleteArticle(id int) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db: db}
}

func (a *articleRepository) CreateArticle(article *model.Articles) error {
	return a.db.Create(article).Error
}

func (a *articleRepository) GetArticleById(id int) (*model.Articles, error) {
	var article *model.Articles
	if err := a.db.Take(&article, id).Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (a *articleRepository) GetAllArticles() ([]model.Articles, error) {
	var users []model.Articles

	if err := a.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (a *articleRepository) UpdateArticle(article *model.Articles) error {
	return a.db.Where("id = ?", article.Id).Updates(&article).Error

}

func (a *articleRepository) DeleteArticle(id int) error {
	var article model.Articles
	return a.db.Table(article.TableName()).Delete(nil, id).Error
}

package repository

import (
	"ArticleVista/internal/common"
	"ArticleVista/internal/model"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	CreateArticle(article *model.Articles) error
	GetArticleById(id int) (*model.Articles, error)
	GetAllArticles(options GetArticleOptions) ([]model.Articles, error)
	UpdateArticle(article *model.Articles) error
	DeleteArticle(id int) error
	GetTotalDocument() (int64, error)
}

type GetArticleOptions struct {
	Filter      map[string]interface{}
	Pagination  common.Pagination
	Sort        string   //createAt desc
	SelectField []string //title content
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

var defaultLimit int = 20

func (a *articleRepository) GetAllArticles(options GetArticleOptions) ([]model.Articles, error) {
	var articles []model.Articles

	query := a.db.Model(&model.Articles{})

	//filter
	for key, value := range options.Filter {
		query = query.Where(key+"= ?", value)
	}
	//Select Field
	if len(options.SelectField) > 0 {
		query = query.Select(options.SelectField)
	}
	//pagination
	if options.Pagination.Limit <= 0 {
		options.Pagination.Limit = defaultLimit
	}
	query = query.Limit(options.Pagination.Limit).Offset(options.Pagination.Skip)
	//Sort
	if options.Sort != "" {
		query = query.Order(options.Sort)
	}
	//Loop to find
	if err := query.Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *articleRepository) UpdateArticle(article *model.Articles) error {
	return a.db.Where("id = ?", article.Id).Updates(&article).Error

}

func (a *articleRepository) DeleteArticle(id int) error {
	var article model.Articles
	return a.db.Delete(&article, id).Error
}

func (a *articleRepository) GetTotalDocument() (int64, error) {
	var article model.Articles
	var count int64
	if err := a.db.Model(&article).Count(&count).Error; err != nil {
		return -1, err
	}
	return count, nil
}

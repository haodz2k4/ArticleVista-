package repository

import "ArticleVista/internal/model"

type ArticleResponsitory interface {
	CreateArticle(article *model.Articles) error
	GetArticleById(id int) (*model.Articles, error)
	GetAllArticles() ([]model.Articles, error)
	UpdateArticle(article *model.Articles) error
	DeleteArticle(id int) error
}

package service

import (
	"ArticleVista/internal/model"
	"ArticleVista/internal/repository"
)

type ArticleService interface {
	CreateArticle(article *model.Articles) error
	GetArticleById(id int) (*model.Articles, error)
	GetAllArticles() ([]model.Articles, error)
	UpdateArticle(article *model.Articles) error
	DeleteArticle(id int) error
}
type articleService struct {
	repo repository.ArticleRepository
}

func NewArticleService(repo repository.ArticleRepository) ArticleService {
	return &articleService{repo}
}
func (s *articleService) CreateArticle(article *model.Articles) error {
	return s.repo.CreateArticle(article)
}

func (s *articleService) GetArticleById(id int) (*model.Articles, error) {
	return s.repo.GetArticleById(id)
}

func (s *articleService) GetAllArticles() ([]model.Articles, error) {
	return s.repo.GetAllArticles()
}

func (s *articleService) UpdateArticle(article *model.Articles) error {
	return s.repo.UpdateArticle(article)
}

func (s *articleService) DeleteArticle(id int) error {
	return s.repo.DeleteArticle(id)
}

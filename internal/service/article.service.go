package service

import (
	"ArticleVista/internal/model"
	"ArticleVista/internal/repository"
	"github.com/go-playground/validator/v10"
)

type ArticleService interface {
	CreateArticle(article *model.Articles) error
	GetArticleById(id int) (*model.Articles, error)
	GetAllArticles() ([]model.Articles, error)
	UpdateArticle(article *model.Articles) error
	DeleteArticle(id int) error
}

// Validate
var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateArticle(article *model.Articles) error {
	return validate.Struct(article)
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

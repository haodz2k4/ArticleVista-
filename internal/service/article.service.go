package service

import (
	"ArticleVista/internal/model"
	"ArticleVista/internal/repository"
	"github.com/go-playground/validator/v10"
)

type ArticleService interface {
	CreateArticle(article *model.Articles) error
	GetArticleById(id int) (*model.Articles, error)
	GetAllArticles(option repository.GetArticleOptions) ([]model.Articles, error)
	UpdateArticle(article *model.Articles) error
	DeleteArticle(id int) error
	GetTotalDocument() (int64, error)
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

func (s *articleService) GetAllArticles(option repository.GetArticleOptions) ([]model.Articles, error) {
	return s.repo.GetAllArticles(option)
}

func (s *articleService) UpdateArticle(article *model.Articles) error {
	return s.repo.UpdateArticle(article)
}

func (s *articleService) DeleteArticle(id int) error {
	return s.repo.DeleteArticle(id)
}

func (s *articleService) GetTotalDocument() (int64, error) {
	return s.repo.GetTotalDocument()
}

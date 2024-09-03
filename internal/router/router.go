package router

import (
	"ArticleVista/internal/handler"
	"ArticleVista/internal/repository"
	"ArticleVista/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	articleRepo := repository.NewArticleRepository(db)
	articleService := service.NewArticleService(articleRepo)
	articleHandle := handler.NewArticleHandler(articleService)
	r.Group("/articles")
	{
		r.POST("", articleHandle.CreateArticle)
	}

	return r
}

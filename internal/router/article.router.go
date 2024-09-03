package router

import (
	"ArticleVista/internal/handler"
	"ArticleVista/internal/repository"
	"ArticleVista/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterArticleRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	articleRepo := repository.NewArticleRepository(db)
	articleService := service.NewArticleService(articleRepo)
	articleHandle := handler.NewArticleHandler(articleService)

	rg.POST("", articleHandle.CreateArticle)
	rg.GET("", articleHandle.GetAllArticles)
	rg.PATCH("/:id", articleHandle.UpdateArtilce)
	rg.DELETE("/:id", articleHandle.DeleteArticle)
}

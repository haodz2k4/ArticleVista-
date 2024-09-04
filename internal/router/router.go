package router

import (
	"ArticleVista/internal/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.ErrorHandling())
	articleGroup := r.Group("/articles")
	RegisterArticleRoutes(articleGroup, db)

	return r
}

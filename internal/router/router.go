package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUpRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	articleGroup := r.Group("/articles")
	RegisterArticleRoutes(articleGroup, db)

	return r
}

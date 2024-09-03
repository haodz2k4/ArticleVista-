package handler

import (
	"ArticleVista/internal/model"
	"ArticleVista/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ArticleHandler struct {
	service service.ArticleService
}

func NewArticleHandler(s service.ArticleService) *ArticleHandler {
	return &ArticleHandler{service: s}
}

func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	var article model.Articles
	if err := c.ShouldBind(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cannot be create article", "error": err.Error()})
		return
	}
	if err := h.service.CreateArticle(&article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot be create article", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"article": article})
}

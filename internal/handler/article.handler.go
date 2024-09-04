package handler

import (
	"ArticleVista/internal/common"
	"ArticleVista/internal/model"
	"ArticleVista/internal/repository"
	"ArticleVista/internal/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
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

	if err := service.ValidateArticle(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation_error": err.Error()})
		return
	}

	if err := h.service.CreateArticle(&article); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "cannot be create article", "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"article": article})
}

func (h *ArticleHandler) GetAllArticles(c *gin.Context) {
	var option repository.GetArticleOptions

	if title := c.Query("title"); title != "" {
		if option.Filter == nil {
			option.Filter = make(map[string]interface{})
		}
		option.Filter["title"] = title
	}
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "30"))

	totalDocument, err := h.service.GetTotalDocument()
	if err != nil {
		c.Error(common.NewAppError(http.StatusBadRequest, "cannot be get total: "+err.Error()))
	}
	option.Pagination = *common.NewPagination(page, limit, totalDocument)

	//sorting
	sort := c.Query("sort")
	if sort != "" {
		option.Sort = sort
	}

	//select Fields
	selectField := c.Query("only")
	if selectField != "" {
		option.SelectField = strings.Split(selectField, " ")
	}
	fmt.Println(option)
	articles, err := h.service.GetAllArticles(option)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "cannot get article", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"articles": articles})
}

func (h *ArticleHandler) UpdateArtilce(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.Error(common.NewAppError(http.StatusBadRequest, "invalid cache id"))
		return
	}

	var article model.Articles

	article.Id = id

	if err := c.ShouldBind(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := h.service.UpdateArticle(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"article": article})
}

func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := h.service.DeleteArticle(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Delete article successfully"})
}

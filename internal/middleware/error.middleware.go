package middleware

import (
	"ArticleVista/internal/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandling() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()

		if len(c.Errors) > 0 {

			err := c.Errors.Last().Err

			if appErr, ok := err.(*common.AppError); ok {
				c.JSON(appErr.Code, gin.H{"error": appErr.Message})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server"})
			}
		}

	}
}

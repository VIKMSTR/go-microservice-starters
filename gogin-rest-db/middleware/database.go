package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Database(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("DB", db)
		c.Next()
	}
}

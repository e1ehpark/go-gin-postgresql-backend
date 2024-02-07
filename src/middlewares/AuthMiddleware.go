package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		if apiKey == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized to perform request. Please get a valid API key"})
			return
		}
		c.Next()
	}
}

func RegisterMiddlewares(router *gin.Engine) {
	router.Use(AuthMiddleware())
}
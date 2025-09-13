package Middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	allowedOrigins := []string{"http://localhost:3000", "https://myblog.com", "http://localhost:8080/"}
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		for _, o := range allowedOrigins {
			if o == origin {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			}
		}
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-API-Key, X-Request-ID")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

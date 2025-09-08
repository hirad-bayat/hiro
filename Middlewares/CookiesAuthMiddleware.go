package Middlewares

import (
	"Hiro/Database"
	"Hiro/Models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AuthRequired middleware to protect routes
func CookiesAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		userID := session.Get("user_id")

		if userID == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		// Check if user still exists in database
		var user Models.User
		if err := Database.DB.First(&user, userID).Error; err != nil {
			session.Clear()
			session.Save()
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		c.Set("current_user", user)
		c.Next()
	}
}

package Routes

import (
	"Hiro/Controllers"
	"Hiro/Middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterWebRoutes(r *gin.Engine) {
	// Public routes
	r.GET("/", Controllers.HomePage)
	r.GET("/login", Controllers.LoginPage)
	r.GET("/register", Controllers.RegisterPage)
	r.POST("/login", Controllers.LoginHandler)
	r.POST("/register", Controllers.RegisterHandler)

	// Protected routes (require authentication)
	protected := r.Group("/")
	protected.Use(Middlewares.CookiesAuthMiddleware())
	{
		protected.GET("/dashboard", Controllers.DashboardPage)
		protected.POST("/logout", Controllers.LogoutHandler)
	}
}

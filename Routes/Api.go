package Routes

import (
	"Hiro/Controllers"
	"Hiro/Middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	// Public routes
	r.GET("/users", Controllers.GetUsers)
	r.POST("/users", Controllers.CreateUser)

	// Protected routes
	protected := r.Group("/")
	protected.Use(Middlewares.JWTMiddleware())
	{
		protected.GET("/users/:id", Controllers.GetUser)
		protected.PUT("/users/:id", Controllers.UpdateUser)
		protected.DELETE("/users/:id", Controllers.DeleteUser)
	}
}

func RegisterBlogRoutes(r *gin.Engine) {
	r.GET("/blogs", Controllers.GetBlogs)
	r.POST("/blogs", Controllers.CreateBlog)
	r.GET("/blogs/:id", Controllers.GetBlog)
	r.PUT("/blogs/:id", Controllers.UpdateBlog)
	r.DELETE("/blogs/:id", Controllers.DeleteBlog)
}

// RegisterAuthRoutes adds authentication routes
func RegisterAuthRoutes(r *gin.Engine) {
	r.POST("/login", Controllers.Login)
	r.POST("/logout", Middlewares.JWTMiddleware(), Controllers.Logout)
	r.POST("/register", Controllers.Register)
}

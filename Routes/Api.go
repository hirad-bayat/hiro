package Routes

import (
	"Hiro/Controllers"
	UserController "Hiro/Internal/User/Controllers"
	"Hiro/Middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *UserController.UserHandler, r *gin.Engine) {

	// Use consistent base path
	api := r.Group("/api")
	{
		// Public routes - use handler methods consistently
		api.GET("/users", userHandler.GetUsers)
		api.POST("/users", userHandler.CreateUser)

		// Protected routes
		protected := api.Group("/")
		protected.Use(Middlewares.JWTMiddleware())
		{
			protected.GET("/users/:id", userHandler.GetUser)
			//protected.PUT("/users/:id", userHandler.UpdateUser)
			//protected.DELETE("/users/:id", userHandler.DeleteUser)
		}
	}
}

func RegisterBlogRoutes(r *gin.Engine) {
	r.GET("api/blogs", Controllers.GetBlogs)
	r.POST("api/blogs", Controllers.CreateBlog)
	r.GET("api/blogs/:id", Controllers.GetBlog)
	r.PUT("api/blogs/:id", Controllers.UpdateBlog)
	r.DELETE("api/blogs/:id", Controllers.DeleteBlog)
}

// RegisterAuthRoutes adds authentication routes
func RegisterAuthRoutes(r *gin.Engine) {
	r.POST("api/login", Controllers.Login)
	r.POST("api/logout", Middlewares.JWTMiddleware(), Controllers.Logout)
	r.POST("api/register", Controllers.Register)
}

package Routes

import (
	"Hiro/Controllers"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	r.GET("/users", Controllers.GetUsers)
	r.POST("/users", Controllers.CreateUser)
	r.GET("/users/:id", Controllers.GetUser)
	r.PUT("/users/:id", Controllers.UpdateUser)
	r.DELETE("/users/:id", Controllers.DeleteUser)
}

func RegisterBlogRoutes(r *gin.Engine) {
	r.GET("/blogs", Controllers.GetBlogs)
	r.POST("/blogs", Controllers.CreateBlog)
	r.GET("/blogs/:id", Controllers.GetBlog)
	r.PUT("/blogs/:id", Controllers.UpdateBlog)
	r.DELETE("/blogs/:id", Controllers.DeleteBlog)
}

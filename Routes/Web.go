package Routes

import (
	"Hiro/Controllers"
	"github.com/gin-gonic/gin"
)

func RegisterWebRoutes(r *gin.Engine) {
	r.GET("/", Controllers.GetBlogs)
	r.GET("/contact", Controllers.GetBlog)
}

package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "home", gin.H{
		"title": "Home",
	})
}

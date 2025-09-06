package Controllers

import (
	"Hiro/Database"
	"Hiro/Models"
	"github.com/gin-gonic/gin"
)

func GetBlogs(c *gin.Context) {
	var blogs []Models.Blog
	Database.DB.Preload("User").Find(&blogs) // Preload user relation
	c.JSON(200, blogs)
}

func CreateBlog(c *gin.Context) {
	var blog Models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	Database.DB.Create(&blog)
	Database.DB.Preload("User").First(&blog, blog.ID)

	c.JSON(201, blog)
}

func GetBlog(c *gin.Context) {
	var blog Models.Blog
	id := c.Param("id")
	if err := Database.DB.Preload("User").First(&blog, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Blog not found"})
		return
	}
	c.JSON(200, blog)
}

func UpdateBlog(c *gin.Context) {
	var blog Models.Blog
	id := c.Param("id")
	if err := Database.DB.First(&blog, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Blog not found"})
		return
	}

	var input Models.Blog
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	blog.Title = input.Title
	blog.Content = input.Content
	blog.UserID = input.UserID

	Database.DB.Save(&blog)
	c.JSON(200, blog)
}

func DeleteBlog(c *gin.Context) {
	var blog Models.Blog
	id := c.Param("id")
	if err := Database.DB.First(&blog, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Blog not found"})
		return
	}
	Database.DB.Delete(&blog)
	c.JSON(200, gin.H{"message": "Blog deleted successfully"})
}

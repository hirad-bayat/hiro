package Controllers

import (
	"Hiro/Database"
	"Hiro/Models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []Models.User
	Database.DB.Find(&users)
	c.JSON(200, users)
}

func CreateUser(c *gin.Context) {
	var user Models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	Database.DB.Create(&user)
	c.JSON(201, user)
}

func GetUser(c *gin.Context) {
	var user Models.User
	id := c.Param("id")
	result := Database.DB.First(&user, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {
	var user Models.User
	id := c.Param("id")
	if err := Database.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	var input Models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password
	Database.DB.Save(&user)

	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	var user Models.User
	id := c.Param("id")
	if err := Database.DB.First(&user, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	Database.DB.Delete(&user)
	c.JSON(200, gin.H{"message": "User deleted successfully"})
}

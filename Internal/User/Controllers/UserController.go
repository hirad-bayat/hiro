package Controllers

import (
	"Hiro/Database"
	"Hiro/Internal/User/Services"
	"Hiro/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService Services.UserService
}

func NewUserHandler(userService Services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user Models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := h.userService.CreateUser(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.userService.GetUser(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// @BasePath /api/
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Get users list
// @Router /users/ [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	//user, err := h.userService.GetUser(c.Request.Context(), uint(id))
	user, err := h.userService.GetAllUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User list not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

//func CreateUser(c *gin.Context) {
//	var user Models.User
//	if err := c.ShouldBindJSON(&user); err != nil {
//		c.JSON(400, gin.H{"error": err.Error()})
//		return
//	}
//	Database.DB.Create(&user)
//	c.JSON(201, user)
//}

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

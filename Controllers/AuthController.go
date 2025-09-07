package Controllers

import (
	"Hiro/Database"
	"Hiro/Middlewares"
	"Hiro/Models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login handles user authentication
func Login(c *gin.Context) {
	var loginReq LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Find user by email
	var user Models.User
	result := Database.DB.Where("email = ?", loginReq.Email).First(&user)
	if result.Error != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 24 hours expiration
	})

	tokenString, err := token.SignedString(Middlewares.JwtSecret)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not generate token"})
		return
	}

	// Store token in database
	expiresAt := time.Now().Add(time.Hour * 24)
	jwtToken := Models.AccessToken{
		Token:     tokenString,
		UserID:    user.ID,
		ExpiresAt: expiresAt,
		Revoked:   false,
	}

	Database.DB.Create(&jwtToken)

	c.JSON(200, gin.H{
		"token":      tokenString,
		"expires_at": expiresAt,
	})
}

// Register handles user registration
func Register(c *gin.Context) {
	var user Models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Could not hash password"})
		return
	}
	user.Password = string(hashedPassword)

	// Create user
	result := Database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(201, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}

// Logout revokes a JWT token
func Logout(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	// Revoke token in database
	result := Database.DB.Model(&Models.AccessToken{}).Where("token = ?", tokenString).Update("revoked", true)
	if result.Error != nil || result.RowsAffected == 0 {
		c.JSON(400, gin.H{"error": "Could not logout"})
		return
	}

	c.JSON(200, gin.H{"message": "Logged out successfully"})
}

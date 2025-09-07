package Middlewares

import (
	"Hiro/Database"
	"Hiro/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var JwtSecret = []byte("Hiro") // Change this to a secure secret in production

// AuthMiddleware checks for a valid JWT token
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix if present
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		// Check if token exists in database and is not revoked
		var storedToken Models.AccessToken
		result := Database.DB.Where("token = ? AND revoked = ?", tokenString, false).First(&storedToken)
		if result.Error != nil {
			c.JSON(401, gin.H{"error": "Invalid or revoked token"})
			c.Abort()
			return
		}

		// Check if token is expired
		if time.Now().After(storedToken.ExpiresAt) {
			c.JSON(401, gin.H{"error": "Token expired"})
			c.Abort()
			return
		}

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return JwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.JSON(401, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Extract claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			userID := uint(claims["user_id"].(float64))
			c.Set("userID", userID)
		} else {
			c.JSON(401, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}

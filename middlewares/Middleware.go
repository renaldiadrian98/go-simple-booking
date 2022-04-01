package middlewares

import (
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func MiddlewareToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(403, gin.H{
			"success": false,
			"message": "No Authorization Header",
			"data":    nil,
		})
		return
	}

	arrAuth := strings.Fields(authHeader)
	tokenString := arrAuth[1]

	secretKey := os.Getenv("SECRETKEY")
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	for key, val := range claims {
		if key == "user_id" {
			c.Set("userId", val)
		}
		if key == "email" {
			c.Set("email", val)
		}
	}

	c.Next()
}

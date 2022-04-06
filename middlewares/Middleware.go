package middlewares

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func middlewareParseToken(authHeader string) (jwt.MapClaims, error) {
	arrAuth := strings.Fields(authHeader)
	tokenString := arrAuth[1]

	secretKey := os.Getenv("SECRETKEY")
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	return claims, err
}

func MiddlewareManagerOnly(c *gin.Context) {
	roleId := c.MustGet("role_id")
	roleIdrStr := fmt.Sprint(roleId)
	roleIdInt, _ := strconv.Atoi(roleIdrStr)
	fmt.Println(roleIdInt)
	if roleIdInt != 1 {
		c.AbortWithStatusJSON(403, gin.H{
			"success": false,
			"message": "Unauthorized",
			"data":    nil,
		})
		return
	}
	c.Next()
}

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
	claims, err := middlewareParseToken(authHeader)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	for key, val := range claims {
		// Check expire token
		if key == "expire" {
			if time.Now().Unix() == val {
				c.AbortWithStatusJSON(401, gin.H{
					"success": false,
					"message": err.Error(),
					"data":    nil,
				})
				return
			}
		}
		if key == "user_id" {
			c.Set("userId", val)
		}
		if key == "email" {
			c.Set("email", val)
		}
		if key == "role_id" {
			c.Set("role_id", val)
		}
	}

	c.Next()
}

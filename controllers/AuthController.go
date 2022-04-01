package controllers

import (
	"go-simple-booking/helpers"
	"go-simple-booking/models"
	"go-simple-booking/responses"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func AuthLogin(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	// Check if user exist
	user, isExist := models.CheckExistingEmail(input.Email)
	if isExist == false {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": "Email or password might be wrong",
		})
		return
	}

	// Check if password is wrong
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil || err == bcrypt.ErrMismatchedHashAndPassword {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": "Invalid Password",
		})
		return
	}

	token, err := helpers.GenerateToken(user.ID, user.Email)

	// Fill response
	var userResponse responses.User
	userResponse.Email = user.Email
	userResponse.Token = token
	c.JSON(200, gin.H{
		"success": true,
		"message": "Login Success",
		"data":    userResponse,
	})
	return

}

func AuthRegister(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	// Check if user exist
	_, isExist := models.CheckExistingEmail(input.Email)
	if isExist == true {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": "Email already registered",
		})
		return
	}

	// Generate password
	pass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	// Save to db
	var user models.User
	user.Email = input.Email
	user.Password = string(pass)
	models.DB.Save(&user)

	// Generate token
	token, err := helpers.GenerateToken(user.ID, user.Email)

	var userResponse responses.User
	userResponse.Email = user.Email
	userResponse.Token = token
	c.JSON(200, gin.H{
		"success": true,
		"message": "Account Created",
		"data":    userResponse,
	})
	return

}

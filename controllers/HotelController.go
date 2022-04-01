package controllers

import (
	"go-simple-booking/helpers"
	"go-simple-booking/models"

	"github.com/gin-gonic/gin"
)

func HotelStore(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	userIdInt := helpers.HelpersGetUserIdInt(c)

	var hotel models.Hotel
	hotel.Name = input.Name
	hotel.UserId = userIdInt
	err := models.DB.Save(&hotel).Error
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": "Failed Saving Hotel",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Success",
		"data":    hotel,
	})
	return
}

func HotelUpdate(c *gin.Context) {
	// Get input
	var input struct {
		Name string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	// Update hotel

}

func HotelGet(c *gin.Context) {
	pageStr := c.Query("page")
	page, err := helpers.HelpersPageQueryToInt(pageStr)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	hotel, err := models.HotelGet(page)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	paginationData, err := models.HelpersPaginationData(page, "hotels")
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"success":         true,
		"message":         "success",
		"data":            hotel,
		"pagination_data": paginationData,
	})
	return
}

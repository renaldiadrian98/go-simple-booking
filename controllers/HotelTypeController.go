package controllers

import (
	"go-simple-booking/helpers"
	"go-simple-booking/models"

	"github.com/gin-gonic/gin"
)

func HotelTypeGet(c *gin.Context) {
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

	hotelType, err := models.HotelTypeGet(page)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	paginationData, err := models.HelpersPaginationData(page, "hotel_types")
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
		"message":         "Success",
		"data":            hotelType,
		"pagination_data": paginationData,
	})
	return

}

func HotelTypeStore(c *gin.Context) {
	var input struct {
		HotelId int    `json:"hotel_id" binding:"required"`
		Name    string `json:"name" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	var hotelType models.HotelType
	hotelType.HotelId = input.HotelId
	hotelType.Name = input.Name
	err := models.DB.Save(&hotelType).Error
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": "Failed Saving Hotel Type",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Success",
		"data":    hotelType,
	})
	return
}

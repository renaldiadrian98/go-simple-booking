package controllers

import (
	"go-simple-booking/models"

	"github.com/gin-gonic/gin"
)

func HotelRoomStore(c *gin.Context) {
	var input struct {
		HotelId     int `json:"hotel_id" binding:"required"`
		HotelTypeId int `json:"hotel_types_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	var hotelRoom models.HotelRoom
	hotelRoom.HotelId = input.HotelId
	hotelRoom.HotelTypeId = input.HotelTypeId
	err := models.DB.Save(&hotelRoom).Error
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": "Failed Saving Hotel Room",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Success",
		"data":    hotelRoom,
	})
	return
}

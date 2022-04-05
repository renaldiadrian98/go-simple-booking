package controllers

import (
	"go-simple-booking/helpers"
	"go-simple-booking/models"

	"github.com/gin-gonic/gin"
)

func TransactionGet(c *gin.Context) {
	userIdInt := helpers.HelpersGetUserIdInt(c)
	page, err := helpers.HelpersPageQueryToInt(c.Query("page"))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	transaction, err := models.TransactionGet(page, userIdInt)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	paginationData, err := models.HelpersPaginationData(page, "transactions")
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
		"data":            transaction,
		"pagination_data": paginationData,
	})
	return
}

func TransactionStore(c *gin.Context) {
	// Get input
	var input struct {
		HotelRoomId  int    `json:"hotel_room_id" binding:"required"`
		CheckinDate  string `json:"checkin_date" binding:"required"`
		CheckoutDate string `json:"checkout_date" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	// Check if hotel_room_id && checkin_date is available && is_paid == true
	isAvailable := models.TransactionCheckAvailability(input.HotelRoomId, input.CheckinDate, input.CheckoutDate)
	if isAvailable == false {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": "Room Not Available",
			"data":    nil,
		})
		return
	}

	// Save Transaction
	userIdInt := helpers.HelpersGetUserIdInt(c)

	transaction, err := models.TransactionStore(userIdInt, input.HotelRoomId, input.CheckinDate, input.CheckoutDate)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"message": "Success",
		"data":    transaction,
	})
	return
}

func TransactionUpdate(c *gin.Context) {
	// Get input
	var input struct {
		StatusCode    int `json:"status_code" binding:"required"`
		PaidPrice     int `json:"paid_price"`
		TransactionId int `json:"transaction_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	// Update transaction
	transaction, err := models.TransactionUpdate(input.TransactionId, input.StatusCode, input.PaidPrice)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"success": false,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"message": "Success",
		"data":    transaction,
	})
	return
}

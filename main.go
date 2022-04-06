package main

import (
	"fmt"
	"go-simple-booking/controllers"
	"go-simple-booking/middlewares"
	"go-simple-booking/models"

	"github.com/gin-gonic/gin"
)

func main() {
	SetupServer().Run()
	fmt.Println("Server is running")
}

func SetupServer() *gin.Engine {
	r := gin.Default()
	models.Connect()

	// Set views folder
	r.LoadHTMLGlob("views/*")
	r.GET("/", controllers.IndexHomePage)

	// Authentication
	r.POST("/api/register", controllers.AuthRegister)
	r.POST("/api/login", controllers.AuthLogin)

	// Hotel
	r.GET("/api/hotel", controllers.HotelGet)
	r.POST("/api/hotel", middlewares.MiddlewareToken, middlewares.MiddlewareManagerOnly, controllers.HotelStore)
	r.PATCH("/api/hotel/:hotel_id", middlewares.MiddlewareToken, middlewares.MiddlewareManagerOnly, controllers.HotelUpdate)

	// Hotel Type
	r.POST("/api/hotel_type", middlewares.MiddlewareToken, controllers.HotelTypeStore)
	r.GET("/api/hotel_type", controllers.HotelTypeGet)

	// Transaction
	r.GET("/api/transaction", middlewares.MiddlewareToken, controllers.TransactionGet)
	r.POST("/api/transaction", middlewares.MiddlewareToken, controllers.TransactionStore)
	r.POST("/api/transaction/callback", middlewares.MiddlewareToken, controllers.TransactionUpdate)

	return r
}

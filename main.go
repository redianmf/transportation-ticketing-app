package main

import (
	"github.com/gin-gonic/gin"
	"github.com/redianmf/transportation-ticketing-app/controllers"
	"github.com/redianmf/transportation-ticketing-app/initializers"
	"github.com/redianmf/transportation-ticketing-app/middlewares"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectDatabasePg()
}

func main() {
	defer initializers.DB.Close()

	router := gin.Default()

	// Users
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	// Wallets
	router.GET("/wallet", middlewares.ValidateAuth, controllers.GetWalletByUserId)
	router.PATCH("/wallet", middlewares.ValidateAuth, controllers.UpdateWalletByUserId)

	// Transportation Modes
	router.GET("/transportation-modes", middlewares.ValidateAuth, controllers.GetAllTransportationMode)
	router.GET("/transportation-modes/:id", middlewares.ValidateAuth, controllers.GetTransportationModeById)
	router.POST("/transportation-modes", middlewares.ValidateAuth, controllers.InsertTransportationMode)
	router.PATCH("/transportation-modes/:id", middlewares.ValidateAuth, controllers.UpdateTransportationMode)
	router.DELETE("/transportation-modes/:id", middlewares.ValidateAuth, controllers.DeleteTransportationMode)

	router.Run("localhost:8080")
}

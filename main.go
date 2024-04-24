package main

import (
	"github.com/gin-gonic/gin"
	"github.com/redianmf/transportation-ticketing-app/controllers"
	"github.com/redianmf/transportation-ticketing-app/initializers"
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

	router.Run(":8080")
}

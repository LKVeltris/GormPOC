package main

import (
	"myapp/controllers"
	"myapp/database"
	"myapp/repositories"

	"github.com/gin-gonic/gin"
)

func main() {
	// dbTest.DBTest()

	database.Connect()
	router := gin.Default()
	userRepo := new(repositories.UserRepository)
	billRepo := new(repositories.BillsRepository)

	userController := controllers.NewUserController(userRepo)
	billsController := controllers.NewBillsController(billRepo)

	controllers.RegisterUserRoutes(router, userController)
	controllers.RegisterBillingRoutes(router, billsController)

	router.Run(":8080")
}

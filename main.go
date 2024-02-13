package main

import (
	"github.com/e1ehpark/go-gin-postgresql-backend/src/middlewares"
	"github.com/e1ehpark/go-gin-postgresql-backend/src/utils"
	"github.com/e1ehpark/go-gin-postgresql-backend/src/models"
	"github.com/e1ehpark/go-gin-postgresql-backend/src/routes"
)

func main() {
	utils.LoadEnv()
	models.OpenDatabaseConnection()
	models.AutoMigrateModels()
	router := routes.SetupRoutes()
	middlewares.RegisterMiddlewares(router)
	router.Run(":8080")
}
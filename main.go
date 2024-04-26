package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/configs"
	"main.go/middlewares"
	"main.go/router"
)

var db *gorm.DB = configs.SetupDatabaseConnection()

func main() {
	defer configs.CloseDatabaseConnection(db)

	r := gin.Default()

	r.Use(middlewares.CorsMiddleware())

	router.SetupRouter(r, db)

	port := configs.GetPort()

	r.Run(port)
}

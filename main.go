package main

import (
	"latihan-api/app/router"
	"latihan-api/app/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	ginEngine := gin.Default()

	// cors
	ginEngine.Use(utils.Cors)

	router.RoutesApp(ginEngine)

	ginEngine.Run(":8080")
}

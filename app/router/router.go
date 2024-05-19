package router

import (
	"latihan-api/app/controller"
	"latihan-api/app/utils"

	"github.com/gin-gonic/gin"
)

func RoutesApp(router *gin.Engine) {
	user := router.Group("/user")
	{
		user.GET("/list", utils.Cors, controller.List)
		user.POST("/insert", controller.Insert)
		user.PATCH("/update/:id", controller.Update)
		user.DELETE("/delete/:id", controller.Delete)

		user.GET("/external", controller.HitApiExternal)
	}
}

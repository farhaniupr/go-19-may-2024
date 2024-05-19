package controller

import (
	"latihan-api/app/model"
	"latihan-api/app/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {

	var user model.User

	utils.DB.WithContext(c.Request.Context()).Table("user").Select("*").Scan(&user)

	c.JSON(200, map[string]interface{}{
		"DATA": user,
	})
}

func Insert(c *gin.Context) {
	var user model.User

	err := c.ShouldBind(&user)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, map[string]interface{}{
			"ERROR": err.Error(),
		})
		return
	}

	utils.DB.WithContext(c.Request.Context()).Table("user").Create(&user)

	// 201 create
	// 200 OK
	// 404 not found
	// 500 internal server error

	c.JSON(201, map[string]interface{}{
		"DATA": user,
	})

}

func Update(c *gin.Context) {
	var user model.User

	id := c.Param("id")

	err := c.ShouldBind(&user)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, map[string]interface{}{
			"ERROR": err.Error(),
		})
		return
	}

	utils.DB.WithContext(c.Request.Context()).Table("user").Where("id = ?", id).Updates(&user).Scan(&user)

	// 201 create
	// 200 OK
	// 404 not found
	// 500 internal server error

	c.JSON(200, map[string]interface{}{
		"DATA": user,
	})

}

func Delete(c *gin.Context) {
	var user model.User

	id := c.Param("id")

	utils.DB.WithContext(c.Request.Context()).Table("user").Where("id = ?", id).Delete(&user)

	// 201 create
	// 200 OK
	// 404 not found
	// 500 internal server error

	c.JSON(200, map[string]interface{}{
		"DATA": []interface{}{},
	})

}

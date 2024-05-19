package utils

import (
	"github.com/gin-gonic/gin"
)

func Cors(c *gin.Context) {

	if c.Request.Header.Get("Authorization") != "test" {
		c.JSON(401, gin.H{
			"message": "authorization is invalid"})
		c.Abort()
	}

	allowList := map[string]bool{
		"": true,
	}

	if origin := c.Request.Header.Get("Origin"); allowList[origin] {

		c.Header("Access-Control-Allow-Origin", origin)

		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Content-Type", "application/json")

		// post, get , delete, path, put
		if c.Request.Method != "OPTIONS" {
			c.Next()
		} else {
			//optonns
			c.AbortWithStatus(401)
		}
	} else {
		c.JSON(401, gin.H{
			"message": "cors origin not allowed"})
		c.Abort()
	}

	c.Request.Response.Header.Set("Access-Control-Allow-Origin", "*")
	c.Request.Response.Header.Set("Content-Type", "application/json")

}

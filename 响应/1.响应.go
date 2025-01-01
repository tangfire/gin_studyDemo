package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{"code": 0, "msg": "success", "data": gin.H{}})
	})

	r.Run(":8080")
}

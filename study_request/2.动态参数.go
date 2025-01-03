package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	/**
	动态参数
	*/
	r.GET("user/:id", func(c *gin.Context) {
		userId := c.Param("id")

		fmt.Println(userId)
	})

	r.GET("user/:id/:name", func(c *gin.Context) {
		userId := c.Param("id")
		userName := c.Param("name")

		fmt.Println(userId, userName)
	})

	r.Run(":8080")
}

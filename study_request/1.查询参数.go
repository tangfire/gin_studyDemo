package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("", func(c *gin.Context) {
		name := c.Query("name")
		age := c.DefaultQuery("age", "18")
		keyList := c.QueryArray("key")

		fmt.Println(name, age, keyList)
	})

	r.Run(":8080")
}

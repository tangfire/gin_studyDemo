package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("users", func(context *gin.Context) {
		name := context.PostForm("name")
		age, ok := context.GetPostForm("age")

		fmt.Println(name)
		fmt.Println(age, ok)

	})

	r.Run(":8080")
}

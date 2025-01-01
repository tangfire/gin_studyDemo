package main

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func Index(c *gin.Context) {
	c.JSON(200, Response{
		Code:    0,
		Message: "success",
		Data:    map[string]any{},
	})
}

func main() {

	// 1. 初始化
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 2. 挂载路由
	r.GET("/index", Index)

	// 3. 绑定端口,运行
	r.Run("127.0.0.1:8080")
}

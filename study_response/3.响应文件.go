package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	/**
	1.要设置Content-Type,唤起浏览器下载
	2.只能是get请求
	*/
	r.GET("", func(c *gin.Context) {
		c.Header("Content-Type", "application/octet-stream")
		c.Header("Content-Disposition", "attachment; filename=hello123.go")
		c.File("3.响应文件.go")
	})

	r.Run(":8080")
}

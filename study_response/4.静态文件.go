package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	/**
	第一个参数:别名
	第二个参数:实际路径
	*/
	r.Static("st", "static")
	// 静态文件的路径，不能再被路由使用
	r.StaticFile("abc", "static/abc.txt")
	r.Run(":8080")
}

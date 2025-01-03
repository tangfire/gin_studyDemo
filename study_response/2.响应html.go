package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	//r.LoadHTMLGlob("templates/*")
	r.LoadHTMLGlob("templates/index.html")
	r.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", map[string]any{
			"title": "myl",
		})
	})
	r.Run(":8080")
}

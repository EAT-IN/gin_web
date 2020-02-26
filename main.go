package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.Static("/static", "static")
	r.StaticFile("favicon.ico", "template/favicon.ico")

	r.LoadHTMLGlob("template/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1 := r.Group("/v1")
	{
		// 增加事项
		v1.POST("/todo", func(c *gin.Context) {

		})

		// 查看事项
		v1.GET("/todo", func(c *gin.Context) {

		})
		v1.GET("/todo/:id", func(c *gin.Context) {

		})

		// 更改事项
		v1.PUT("/todo/:id", func(c *gin.Context) {

		})

		// 删除事项
		v1.DELETE("/todo/:id", func(c *gin.Context) {

		})

	}

	r.Run(":5000")
}

package router

import (
	"gin_web/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Static("/static", "static")
	r.StaticFile("/favicon.ico", "template/favicon.ico")

	r.LoadHTMLGlob("template/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1 := r.Group("/v1")
	{
		// 增加事项
		v1.POST("/todo", controller.AddTodo)

		// 查看事项
		v1.GET("/todo", controller.SelectTodo)

		// 更改事项
		v1.PUT("/todo/:id", controller.UpdateTodo)

		// 删除事项
		v1.DELETE("/todo/:id", controller.DeleteTodo)

	}

	return r
}

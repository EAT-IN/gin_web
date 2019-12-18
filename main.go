package main

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(200, "index.html", map[string]string{"title": "index"})

}

func main() {

	router := gin.Default()

	// 加载模版文件和静态文件
	router.Static("/static", "./static")
	router.LoadHTMLGlob("./template/*")

	router.GET("/", Index)
	router.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "登录成功!!!",
		})
	})

	router.Run()

}

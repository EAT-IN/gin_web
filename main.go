package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	msg := "小王子"
	c.HTML(200, "index.tmpl", msg)

}

func Home(c *gin.Context) {
	msg := "小王子"
	c.HTML(200, "home.tmpl", msg)

}

func main() {

	router := gin.Default()
	router.LoadHTMLGlob("./templates/*")

	router.GET("/index", Index)
	router.GET("/home", Home)

	if err := router.Run(); err != nil {
		fmt.Println(err.Error())
		return
	}

}

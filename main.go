package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/cookie", func(c *gin.Context) {
		// 先获取cookie值，如果没有就设定一个cookie
		name, err := c.Cookie("name")
		if err != nil {
			name = "not set"
			// name:cookie的名字
			// value：cookie的值
			// maxage：cookie的有效时间， 单位为秒
			// path是指 cookie所在的目录
			// domain 域名
			// secure 是否只能通过https 访问
			// httpone 是否允许别人通过js获取自己的cookie
			c.SetCookie("name", "xmzhang", 3600, "/", "localhost", false, true)
		}
		fmt.Println("cookie的值为：", name)
	})

	r.Run()
}

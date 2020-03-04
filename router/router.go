package router

import (
	"gin_web/model"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

var err error

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
		v1.POST("/todo", func(c *gin.Context) {
			var todo model.Todo
			if err = c.ShouldBind(&todo); err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			} else {
				err = model.DB.Create(&todo).Error
				if err != nil {
					c.JSON(http.StatusOK, gin.H{"error": err.Error()})
					return
				} else {
					c.JSON(http.StatusOK, todo)
				}
			}

		})

		// 查看事项
		v1.GET("/todo", func(c *gin.Context) {
			var todos []*model.Todo
			err = model.DB.Find(&todos).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, todos)

		})
		v1.GET("/todo/:id", func(c *gin.Context) {

		})

		// 更改事项
		v1.PUT("/todo/:id", func(c *gin.Context) {
			id := c.Param("id")
			if id == "" {
				c.JSON(http.StatusOK, gin.H{"error": "id is null"})
			}
			idInt := com.StrTo(id).MustInt()

			todo := model.Todo{
				ID: idInt,
			}
			var status struct {
				Status bool `json:"status"`
			}
			c.ShouldBind(&status)

			err = model.DB.Model(&todo).Update("status", status.Status).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, todo)
		})

		// 删除事项
		v1.DELETE("/todo/:id", func(c *gin.Context) {
			id := c.Param("id")
			if id == "" {
				c.JSON(http.StatusOK, gin.H{"error": "id is null"})
			}
			var todo model.Todo
			model.DB.Where("id=?", id).Delete(&todo)
			c.JSON(http.StatusOK, todo)

		})

	}

	return r
}

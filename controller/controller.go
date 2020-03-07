package controller

import (
	"gin_web/model"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

var err error

func AddTodo(c *gin.Context) {
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

}

func SelectTodo(c *gin.Context) {
	var todos []*model.Todo
	err = model.DB.Find(&todos).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)

}

func UpdateTodo(c *gin.Context) {
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
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusOK, gin.H{"error": "id is null"})
	}
	var todo model.Todo
	model.DB.Where("id=?", id).Delete(&todo)
	c.JSON(http.StatusOK, todo)

}

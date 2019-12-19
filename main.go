package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
)

type Booking struct {
	// 这里定义了一个表单,必须要满足的条件时,不能为空.并且日期需要大于今天 ,这里的time_format 是对前端提交的时间进行格式化,变成time.Time类型
	CheckIn time.Time `form:"check_in" json:"check_in "binding:"required,bookabledate" time_format:"2006-01-02"`
}

func bookableDate(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

func main() {
	route := gin.Default()

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", bookableDate)
	}

	route.GET("/bookable", getBookable)
	route.Run(":8085")
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindQuery(&b); err == nil {
		c.JSON(http.StatusOK, b)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

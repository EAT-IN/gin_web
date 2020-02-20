package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB  *gorm.DB
	err error
)

type Demo struct {
	gorm.Model
	Name string
	Age  int
	High int
}

func InitDB() {
	DB, err = gorm.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	// 关闭表明复数
	DB.SingularTable(true)
	// 开启数据库调试
	DB.LogMode(true)
	//设置最大闲置数量
	DB.DB().SetMaxIdleConns(5)
	// 设置最大连接数
	DB.DB().SetMaxOpenConns(10)
	// 自动建表
	DB.AutoMigrate(&Demo{}) // 自动创建表

}

func main() {
	InitDB()
	d := &Demo{
		Model: gorm.Model{ID: 1},
		Name:  "xmzhang",
		Age:   24,
		High:  170,
	}
	fmt.Println(DB.NewRecord(d))
	// 增
	DB.Create(d)
	fmt.Println(DB.NewRecord(d))

	// 改
	var a Demo
	DB.Model(&a).Updates(map[string]interface{}{
		"name": "111",
		"age":  100,
		"high": 199,
	})

	//删, 逻辑删除，看delete_at的时间就知道了
	DB.Where("id=?", 1).Delete(&a)
	// 真实删除的话，需要加上Unscoped()
	DB.Unscoped().Where("id>?", 1).Delete(&a)

	//查. 如果想查询到逻辑删除的数据 加上Unscoped()
	var data []Demo
	DB.Unscoped().Find(&data)
	for _, v := range data {
		fmt.Println(v.CreatedAt)
	}

}

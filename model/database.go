package model

import (
	"gin_web/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(config.Configs.Key("database_drive").String(), config.Configs.Key("database_dns").String())
	if err != nil {
		panic(err)
	}
	// 关闭表明复数
	DB.SingularTable(true)
	// 开启数据库调试
	DB.LogMode(config.Configs.Key("db_mode").MustBool(true))
	//设置最大闲置数量
	DB.DB().SetMaxIdleConns(5)
	// 设置最大连接数
	DB.DB().SetMaxOpenConns(10)
	// 自动建表
	DB.AutoMigrate(&Todo{}) // 自动创建表go

}

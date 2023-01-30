package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB = Init()

func Init() *gorm.DB {
	dsn := "Lycopene:MiMaJiuShi123321!@tcp(127.0.0.1:3306)/cloud-disk?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return db
}

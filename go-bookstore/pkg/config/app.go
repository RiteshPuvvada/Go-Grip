package config

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql", "root:ritesh@tcp(127.0.0.1:3306)/go?charset=utf8&paraseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	db =d
}

func GetDB() *gorm.DB{
	return db
}
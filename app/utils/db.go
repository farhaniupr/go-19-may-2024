package utils

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	DB *gorm.DB
)

func init() {

	dsn := "root:pass@tcp(127.0.0.1:3306)/latihan?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		log.Println("error :", err.Error())
	} else {
		log.Println("success connect")
	}

	DB = db
}

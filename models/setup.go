package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:admin@tcp(localhost:3306)/asia_quest"))
	if err != nil {
		panic(err)
	}

	err = database.AutoMigrate(&User{}, &Books{}, &Categories{}, BookCategory{})
	if err != nil {
		log.Println(err)
		return
	}

	DB = database
}

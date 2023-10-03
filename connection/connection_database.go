package connection

import (
	"asis_quest/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type connectionDB struct {
	DB *gorm.DB
}

type DatabaseMain interface {
	ConnectDatabase() (*gorm.DB, error)
}

func NewDatabase(DB *gorm.DB) DatabaseMain {
	return &connectionDB{DB: DB}
}

func (db *connectionDB) ConnectDatabase() (*gorm.DB, error) {
	database, err := gorm.Open(mysql.Open("root:admin@tcp(localhost:3306)/asia_quest"))
	if err != nil {
		panic(err)
	}

	err = database.AutoMigrate(&models.User{}, &models.Books{}, &models.Categories{}, models.BookCategory{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return database, nil
}

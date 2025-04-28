package database

import (
	"crud-api-with-go-gin-gorm/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

)

var DB *gorm.DB

func Connect() error{
	var err error
	
	dsn := "root:@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		return err
	}

	err = DB.AutoMigrate(&models.Book{})
	if err != nil {
		return err
	}

	
	
	return nil
}
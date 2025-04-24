package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"crud-api-with-go-gin-gorm/database"
)

func main(){

	err := database.Connect()
	if err != nil{
		log.Fatal("Failed to connect to database", err)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context){
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome To BookStore",
		})
	})

	router.Run()

}
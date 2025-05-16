package controllers

import (
	"crud-api-with-go-gin-gorm/database"
	"crud-api-with-go-gin-gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBook(c *gin.Context){
	id := c.Param("id")
	var book models.Book

	result := database.DB.First(&book, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Book Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, book)

}


func GetAllBooks(c *gin.Context){
	var books []models.Book

	result := database.DB.Find(&books)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch books",
		})
		return
	}
	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request data",
		})
		return
	}

	result := database.DB.Create(&book)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create book",
		})
	}

	c.JSON(http.StatusCreated, book)
}
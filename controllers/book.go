package controllers

import (
	"crud-api-with-go-gin-gorm/database"
	"crud-api-with-go-gin-gorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
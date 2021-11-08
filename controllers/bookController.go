package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hyperyuri/webapi-with-go/database"
	"github.com/hyperyuri/webapi-with-go/models"
)

func ShowBook(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "ID has to be integer",
		})
		return
	}

	db := database.GetDataBase()

	var book models.Book

	err = db.First(&book, newId).Error
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Cannot find book" + err.Error(),
		})
		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDataBase()

	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Cannot bind JSON" + err.Error(),
		})
		return
	}

	err = db.Create(&book).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Cannot created book" + err.Error(),
		})
		return
	}
	c.JSON(200, book)
}

func ShowBooks(c *gin.Context) {
	db := database.GetDataBase()

	var books []models.Book

	err := db.Find(&books).Error

	if err != nil {
		c.JSON(400, gin.H{
			"Error": "Cannot list books" + err.Error(),
		})
		return
	}
	c.JSON(200, books)
}

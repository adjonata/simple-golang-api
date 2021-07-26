package controllers

import (
	"github.com/gin-gonic/gin"
	"go-api/dto"
	"go-api/models"
	"net/http"
)

// FindBooks godoc
// @Summary List all Books
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Book
// @Router /books/ [get]
func FindBooks(c *gin.Context) {
	var books []models.Book

	models.DB.Find(&books)

	c.JSON(http.StatusOK, books)
}

// CreateBook godoc
// @Summary Create new Book
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Book
// @Router /books/ [post]
// @Param requestBody body dto.CreateBookInput true "Body for create book"
func CreateBook(c *gin.Context) {
	var input dto.CreateBookInput

	// Validation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	c.JSON(http.StatusOK, book)
}

// FindBook godoc
// @Summary Find book
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Book
// @Param id path int true "Book ID"
// @Router /books/{id} [get]
func FindBook(c *gin.Context) {
	var book models.Book

	// Find and Check Error
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// UpdateBook godoc
// @Summary Update book
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Book
// @Param id path int true "Book ID"
// @Router /books/{id} [put]
// @Param requestBody body dto.UpdateBookInput true "Body for update book"
func UpdateBook(c *gin.Context) {
	var book models.Book

	// Find and Check Error
	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate Input
	var input dto.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)

	c.JSON(http.StatusOK, book)
}

// DeleteBook godoc
// @Summary Delete book
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"status": true})
}

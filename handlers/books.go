package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/santosh/gingo/db"
	"github.com/santosh/gingo/logger"
	"github.com/santosh/gingo/models"
	"go.uber.org/zap"
)

var zlog *zap.Logger

func init() {
	zlog = logger.Log.With(
		zap.String("package", "handlers"),
	)
}

// GetBooks		 godoc
// @Summary      Get books array
// @Description  Responds with the list of all books as JSON.
// @Tags         books
// @Produce      json
// @Success      200  {array}  models.Book
// @Router       /books [get]
func GetBooks(c *gin.Context) {
	var books []models.Book

	if result := db.DB.Find(&books); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	zlog.Info("GETting all books")

	c.JSON(http.StatusOK, &books)
}

// PostBook		 godoc
// @Summary      Store a new book
// @Description  Takes a book JSON and store in DB. Return saved JSON.
// @Tags         books
// @Produce      json
// @Param        book  body      models.Book  true  "Book JSON"
// @Success      200   {object}  models.Book
// @Router       /books [post]
func PostBook(c *gin.Context) {
	var newBook models.Book

	// Call BindJSON to bind the received JSON to
	// newBook.
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	if result := db.DB.Create(&newBook); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	zlog.Info(fmt.Sprintf("POSTing new book with ISBN: %s", newBook.ISBN))

	c.JSON(http.StatusCreated, &newBook)
}

// GetBookByISBN		 godoc
// @Summary      Get single book by isbn
// @Description  Returns the book whose ISBN value matches the isbn.
// @Tags         books
// @Produce      json
// @Param        isbn  path      string  true  "search book by isbn"
// @Success      200  {object}  models.Book
// @Router       /books/{isbn} [get]
func GetBookByISBN(c *gin.Context) {
	var book models.Book

	if err := db.DB.Where("isbn = ?", c.Param("isbn")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	zlog.Info(fmt.Sprintf("GET request for book: %s", book.ISBN))

	c.JSON(http.StatusOK, &book)
}

// DeleteBookByISBN		 godoc
// @Summary      Remove single book by isbn
// @Description  Delete a single entry from the database based on isbn.
// @Tags         books
// @Produce      json
// @Param        isbn  path      string  true  "delete book by isbn"
// @Success      204
// @Router       /books/{isbn} [delete]
func DeleteBookByISBN(c *gin.Context) {
	id := c.Param("isbn")

	if result := db.DB.Delete(&models.Book{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	zlog.Info(fmt.Sprintf("DELETE request for book: %s", id))

	c.Status(http.StatusNoContent)
}

// UpdateBookByISBN		 godoc
// @Summary      Update single book by isbn
// @Description  Updates and returns a single book whose ISBN value matches the isbn. New data must be passed in the body.
// @Tags         books
// @Produce      json
// @Param        isbn  path      string  true  "update book by isbn"
// @Success      200  {object}  models.Book
// @Router       /books/{isbn} [put]
func UpdateBookByISBN(c *gin.Context) {
	// Get model if exist
	var book models.Book
	var bookUpdate models.Book

	// pull the specific book intry in &book
	if err := db.DB.Where("isbn = ?", c.Param("isbn")).First(&book).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	// get new data from body
	if err := c.ShouldBindJSON(&bookUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	zlog.Info(fmt.Sprintf("PUT request for book: %s", book.ISBN))

	// update and return new body
	db.DB.Model(&book).Where("isbn = ?", c.Param("isbn")).Updates(&bookUpdate)
	c.JSON(http.StatusOK, &book)
}

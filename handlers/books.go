package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/santosh/gingo/db"
	"github.com/santosh/gingo/models"
)

// GetBooks responds with the list of all books as JSON.
func GetBooks(c *gin.Context) {
	c.JSON(http.StatusOK, db.Books)
}

// PostBook takes a book JSON and store in DB.
func PostBook(c *gin.Context) {
	var newBook models.Book

	// Call BindJSON to bind the received JSON to
	// newBook.
	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// Add the new book to the slice.
	db.Books = append(db.Books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

// GetBookByISBN locates the book whose ISBN value matches the isbn
func GetBookByISBN(c *gin.Context) {
	isbn := c.Param("isbn")

	// Loop over the list of books, look for
	// an book whose ISBN value matches the parameter.
	for _, a := range db.Books {
		if a.ISBN == isbn {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}

// func DeleteBookByISBN(c *gin.Context) {}

// func UpdateBookByISBN(c *gin.Context) {}


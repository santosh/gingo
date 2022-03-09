package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/santosh/gingo/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/books", handlers.GetBooks)
	router.GET("/books/:isbn", handlers.GetBookByISBN)
	// router.DELETE("/books/:isbn", handlers.DeleteBookByISBN)
	// router.PUT("/books/:isbn", handlers.UpdateBookByISBN)
	router.POST("/books", handlers.PostBook)

	return router
}

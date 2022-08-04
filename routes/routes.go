package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/santosh/gingo/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/books", handlers.PostBook)
		v1.GET("/books", handlers.GetBooks)
		v1.GET("/books/:isbn", handlers.GetBookByISBN)
		v1.DELETE("/books/:isbn", handlers.DeleteBookByISBN)
		v1.PUT("/books/:isbn", handlers.UpdateBookByISBN)
	}

	return router
}

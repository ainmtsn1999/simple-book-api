package routers

import (
	"github.com/ainmtsn1999/simple-book-api/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {

	router := gin.Default()

	router.GET("/books/:bookId", controllers.GetBook)
	router.GET("/books", controllers.GetAllBooks)
	router.POST("/books", controllers.CreateBook)
	router.PUT("/books/:bookId", controllers.UpdateBook)
	router.DELETE("/books/:bookId", controllers.DeleteBook)

	return router
}

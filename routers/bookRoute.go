package routers

import (
	"challenge-06/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/books/", controllers.CreateBook)
	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:id", controllers.GetBookById)
	router.PUT("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	return router
}

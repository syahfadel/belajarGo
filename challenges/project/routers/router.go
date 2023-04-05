package routers

import (
	"project/controllers"
	"project/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartServer(db *gorm.DB) *gin.Engine {

	bookService := services.BookService{
		DB: db,
	}

	bookController := controllers.BookController{
		DB:          db,
		BookService: &bookService,
	}

	app := gin.Default()
	app.GET("books/", bookController.GetAllBooks)
	app.GET("/books/:id", bookController.GetBookById)
	app.POST("/books/", bookController.CreateBook)
	app.PUT("/books/:id", bookController.UpdateBook)
	app.DELETE("/books/:id", bookController.Delete)
	return app
}

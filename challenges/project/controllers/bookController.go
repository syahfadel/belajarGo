package controllers

import (
	"fmt"
	"net/http"
	"project/entity"
	"project/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookController struct {
	DB          *gorm.DB
	BookService *services.BookService
}

type BookCreateDto struct {
	NameBook string `json:"name_book"`
	Author   string `json:"author"`
}

func (bc *BookController) GetAllBooks(ctx *gin.Context) {
	var (
		books []entity.Book
		err   error
	)

	books, err = bc.BookService.GetAll()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "unable to find data",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}

func (bc *BookController) GetBookById(ctx *gin.Context) {
	var book entity.Book
	rawId := ctx.Param("id")
	id, err := strconv.Atoi(rawId)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Wrong Parameter",
			"error_message": fmt.Sprintf("%s not an integer", rawId),
		})
		return
	}

	book, err = bc.BookService.GetById(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Data with id %v not found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, book)

}

func (bc *BookController) CreateBook(ctx *gin.Context) {
	var bookCreateDto BookCreateDto

	if err := ctx.ShouldBindJSON(&bookCreateDto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	book := entity.Book{
		NameBook: bookCreateDto.NameBook,
		Author:   bookCreateDto.Author,
	}
	result, err := bc.BookService.Create(book)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (bc *BookController) UpdateBook(ctx *gin.Context) {
	var bookCreateDto BookCreateDto
	rawId := ctx.Param("id")
	id, err := strconv.Atoi(rawId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Wrong Parameter",
			"error_message": fmt.Sprintf("%s not an integer", rawId),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&bookCreateDto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	book := entity.Book{
		ID:       uint(id),
		NameBook: bookCreateDto.NameBook,
		Author:   bookCreateDto.Author,
	}

	result, err := bc.BookService.Update(book)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (bc *BookController) Delete(ctx *gin.Context) {
	rawId := ctx.Param("id")
	id, err := strconv.Atoi(rawId)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Wrong Parameter",
			"error_message": fmt.Sprintf("%s not an integer", rawId),
		})
		return
	}

	success := bc.BookService.Delete(id)

	if !success {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Data with id %v not found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Book deleted successfully",
	})
}

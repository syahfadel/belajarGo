package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var BookDatas = []Book{
	{Id: "1", Title: "Golang", Author: "Gopher", Desc: "A book for Go"},
}

func GetAllBooks(ctx *gin.Context) {
	fmt.Println(BookDatas)
	ctx.JSON(http.StatusOK, BookDatas)
}

func GetBookById(ctx *gin.Context) {
	bookId := ctx.Param("id")
	condition := false
	var bookData Book

	for i, book := range BookDatas {
		if bookId == book.Id {
			bookData = BookDatas[i]
			condition = true
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with id %v not found", bookId),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": bookData,
	})
}

func AddBook(ctx *gin.Context) {
	var newBook Book
	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.Id = fmt.Sprintf("%v", len(BookDatas)+1)
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusOK, "Created")
}

func UpdateBook(ctx *gin.Context) {
	bookId := ctx.Param("id")
	condition := false
	var updateBook Book

	if err := ctx.ShouldBindJSON(&updateBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookDatas {
		if bookId == book.Id {
			condition = true
			BookDatas[i] = updateBook
			BookDatas[i].Id = bookId
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Data Not Found",
			"error_message": "Book with id %v not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, "Updated")
}

func DeleteBook(ctx *gin.Context) {
	bookId := ctx.Param("id")
	condition := false
	var bookIndex int

	for i, book := range BookDatas {
		if bookId == book.Id {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Book with id %v not found", bookId),
		})
		return
	}

	copy(BookDatas[bookIndex:], BookDatas[bookIndex+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, "Deleted")

}

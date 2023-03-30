package controllers

import (
	"challenges8/models"
	"challenges8/service"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.GetAllBooks())
}

func GetBookById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Wrong Parameter",
			"error_message": fmt.Sprintf("%v not an integer", id),
		})
		return
	}

	var book models.Book
	book, success := service.GetBookById(id)
	if !success {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Data with id %v not found", id),
		})
		return
	}
	ctx.JSON(http.StatusOK, book)
}

func CreateBook(ctx *gin.Context) {
	var newBook models.Book
	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	success := service.CreateBook(newBook)
	if !success {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, "Created")
}

func UpdateBook(ctx *gin.Context) {
	var updateBook models.Book
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Wrong Parameter",
			"error_message": fmt.Sprintf("%v not an integer", id),
		})
		return
	}

	err = ctx.ShouldBindJSON(&updateBook)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	updateBook.Id = id

	success := service.UpdateBook(updateBook)
	if !success {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("Data with id %v updated", id))
}

func DeleteBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Wrong Parameter",
			"error_message": fmt.Sprintf("%v not an integer", id),
		})
		return
	}

	success := service.DeleteBook(id)

	if !success {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("Data with id %v not found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, fmt.Sprintf("Data with id %v deleted", id))

}

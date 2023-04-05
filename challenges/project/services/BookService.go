package services

import (
	"project/entity"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type BookService struct {
	DB *gorm.DB
}

func (bs *BookService) GetAll() ([]entity.Book, error) {
	var books []entity.Book
	if err := bs.DB.Find(&books).Error; err != nil {
		return books, err
	}
	return books, nil
}

func (bs *BookService) GetById(id int) (entity.Book, error) {
	var book entity.Book
	if err := bs.DB.First(&book, "id = ?", id).Error; err != nil {
		return book, err
	}

	return book, nil
}

func (bs *BookService) Create(book entity.Book) (entity.Book, error) {
	if err := bs.DB.Create(&book).Error; err != nil {
		return entity.Book{}, err
	}

	return book, nil
}

func (bs *BookService) Update(book entity.Book) (entity.Book, error) {
	var updatedBook entity.Book
	if err := bs.DB.Model(&updatedBook).Clauses(clause.Returning{}).Where("id = ?", book.ID).Updates(&book).Error; err != nil {
		return entity.Book{}, err
	}

	return updatedBook, nil
}

func (bs *BookService) Delete(id int) bool {
	err := bs.DB.Where("id = ?", id).Delete(entity.Book{}).Error

	if err != nil {
		return false
	}

	return true
}

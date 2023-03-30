package service

import (
	"challenges8/connection"
	"challenges8/models"
	"fmt"

	_ "github.com/lib/pq"
)

func GetAllBooks() []models.Book {
	var books = []models.Book{}
	sqlStatement := `SELECT * FROM bookdatas`
	rows, err := connection.Db.Query(sqlStatement)
	if err != nil {
		fmt.Println("FAILED AT QUERY")
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var book = models.Book{}
		err = rows.Scan(&book.Id, &book.Title, &book.Author, &book.Desc)
		if err != nil {
			fmt.Println("FAILED AT SCAN")
			panic(err)
		}

		books = append(books, book)
	}

	return books
}

func GetBookById(id int) (models.Book, bool) {
	var book models.Book
	sqlStatement := `
	SELECT * FROM bookdatas
	WHERE id = $1`

	err := connection.Db.QueryRow(sqlStatement, id).
		Scan(&book.Id, &book.Title, &book.Author, &book.Desc)
	if err != nil {
		return book, false
	}

	return book, true
}

func CreateBook(newBook models.Book) bool {
	sqlStatement := `
	INSERT INTO bookdatas (title, author, des)
	VALUES
	($1, $2, $3)
	`
	_, err := connection.Db.Exec(sqlStatement, newBook.Title, newBook.Author, newBook.Desc)
	if err != nil {
		return false
	}

	return true
}

func UpdateBook(updateBook models.Book) bool {

	sqlStatement := fmt.Sprintf(`
	UPDATE bookdatas 
	SET id = %v`, updateBook.Id)

	if updateBook.Title != "" {
		sqlStatement += fmt.Sprintf(`, title = '%s'`, updateBook.Title)
	}
	if updateBook.Author != "" {
		sqlStatement += fmt.Sprintf(`, author = '%s'`, updateBook.Author)
	}
	if updateBook.Desc != "" {
		sqlStatement += fmt.Sprintf(`, des = '%s'`, updateBook.Desc)
	}

	sqlStatement += fmt.Sprintf("WHERE id = %v", updateBook.Id)
	_, err := connection.Db.Exec(sqlStatement)
	if err != nil {
		return false
	}
	return true
}

func DeleteBook(id int) bool {
	sqlStatement := `
	DELETE FROM bookdatas
	WHERE id = $1
	`

	_, err := connection.Db.Exec(sqlStatement, id)
	if err != nil {
		return false
	}
	return true

}

package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/database"
	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/middlewares"
	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/models"
	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/repositories"
	"github.com/labstack/echo/v4"
)

func GetAllBook(c echo.Context) error {
	books, err := repositories.GetAllBook(database.DbConnection)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all books",
		"data":    books,
	})
}

func CreateBook(c echo.Context) error {
	book := models.Book{}
	if err := c.Bind(&book); err != nil {
		return err
	}
	if err := c.Validate(book); err != nil {
		return err
	}

	data, err := repositories.InsertBook(database.DbConnection, book)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create new book",
		"data":    data,
	})
}

func UpdateBookById(c echo.Context) error {
	book := models.Book{}
	if err := c.Bind(&book); err != nil {
		return err
	}
	if err := c.Validate(book); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	book.Id = id

	err := repositories.UpdateBook(database.DbConnection, book)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success update book by id",
		"data":    book,
	})
}

func DeleteBookById(c echo.Context) error {
	book := models.Book{}
	id, _ := strconv.Atoi(c.Param("id"))
	book.Id = id

	err := repositories.DeleteBook(database.DbConnection, book)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success delete book by id",
		"data":    nil,
	})
}

func BorrowBook(c echo.Context) error {
	book := models.Book{}

	id, _ := strconv.Atoi(c.Param("id"))
	book, err := repositories.GetBookById(database.DbConnection, id)

	if err != nil {
		return err
	}

	if book.IsBorrowed {
		return errors.New("book already borrowed")
	}

	userId := middlewares.ExtractTokenUserId(c)

	book.IsBorrowed = true
	book.BorrowedBy = userId

	err = repositories.UpdateBook(database.DbConnection, book)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success borrowing book",
		"data":    book,
	})
}

func ReturnBook(c echo.Context) error {
	book := models.Book{}

	id, _ := strconv.Atoi(c.Param("id"))
	book, err := repositories.GetBookById(database.DbConnection, id)

	if err != nil {
		return err
	}

	if !book.IsBorrowed {
		return errors.New("book is not borrowed by anyone")
	}

	userId := middlewares.ExtractTokenUserId(c)
	if book.BorrowedBy != userId {
		return errors.New("book is not borrowed by you")
	}

	book.IsBorrowed = false

	err = repositories.UpdateBook(database.DbConnection, book)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success returning book",
		"data":    book,
	})
}

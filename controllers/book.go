package controllers

import (
	"net/http"
	"strconv"

	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/database"
	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/models"
	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/repositories"
	"github.com/labstack/echo/v4"
)

func GetAllBook(c echo.Context) error {
	books, err := repositories.GetAllBook(database.DbConnection)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": http.StatusText(http.StatusInternalServerError),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all books",
		"data":    books,
	})
}

func CreateBook(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	data, err := repositories.InsertBook(database.DbConnection, book)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": http.StatusText(http.StatusInternalServerError),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create new book",
		"data":    data,
	})
}

func UpdateBookById(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)
	id, _ := strconv.Atoi(c.Param("id"))
	book.Id = id

	err := repositories.UpdateBook(database.DbConnection, book)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": http.StatusText(http.StatusInternalServerError),
			"data":    nil,
		})
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
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": http.StatusText(http.StatusInternalServerError),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success delete book by id",
		"data":    nil,
	})
}

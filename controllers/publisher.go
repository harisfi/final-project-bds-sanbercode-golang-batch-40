package controllers

import (
	"net/http"
	"strconv"

	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/database"
	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/models"
	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/repositories"
	"github.com/labstack/echo/v4"
)

func GetAllPublisher(c echo.Context) error {
	publishers, err := repositories.GetAllPublisher(database.DbConnection)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": http.StatusText(http.StatusInternalServerError),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all publishers",
		"data":    publishers,
	})
}

func CreatePublisher(c echo.Context) error {
	publisher := models.Publisher{}
	c.Bind(&publisher)

	data, err := repositories.InsertPublisher(database.DbConnection, publisher)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": http.StatusText(http.StatusInternalServerError),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create new publisher",
		"data":    data,
	})
}

func UpdatePublisherById(c echo.Context) error {
	publisher := models.Publisher{}
	c.Bind(&publisher)
	id, _ := strconv.Atoi(c.Param("id"))
	publisher.Id = id

	err := repositories.UpdatePublisher(database.DbConnection, publisher)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": http.StatusText(http.StatusInternalServerError),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success update publisher by id",
		"data":    publisher,
	})
}

func DeletePublisherById(c echo.Context) error {
	publisher := models.Publisher{}
	id, _ := strconv.Atoi(c.Param("id"))
	publisher.Id = id

	err := repositories.DeletePublisher(database.DbConnection, publisher)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": http.StatusText(http.StatusInternalServerError),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success delete publisher by id",
		"data":    nil,
	})
}

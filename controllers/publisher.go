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
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all publishers",
		"data":    publishers,
	})
}

func CreatePublisher(c echo.Context) error {
	publisher := models.Publisher{}
	if err := c.Bind(&publisher); err != nil {
		return err
	}
	if err := c.Validate(publisher); err != nil {
		return err
	}

	data, err := repositories.InsertPublisher(database.DbConnection, publisher)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create new publisher",
		"data":    data,
	})
}

func UpdatePublisherById(c echo.Context) error {
	publisher := models.Publisher{}
	if err := c.Bind(&publisher); err != nil {
		return err
	}
	if err := c.Validate(publisher); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	publisher.Id = id

	err := repositories.UpdatePublisher(database.DbConnection, publisher)

	if err != nil {
		return err
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
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success delete publisher by id",
		"data":    nil,
	})
}

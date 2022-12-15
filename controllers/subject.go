package controllers

import (
	"net/http"
	"strconv"

	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/database"
	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/models"
	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/repositories"
	"github.com/labstack/echo/v4"
)

func GetAllSubject(c echo.Context) error {
	subjects, err := repositories.GetAllSubject(database.DbConnection)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": http.StatusText(http.StatusInternalServerError),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all subjects",
		"data":    subjects,
	})
}

func CreateSubject(c echo.Context) error {
	subject := models.Subject{}
	c.Bind(&subject)

	data, err := repositories.InsertSubject(database.DbConnection, subject)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": http.StatusText(http.StatusInternalServerError),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create new subject",
		"data":    data,
	})
}

func UpdateSubjectById(c echo.Context) error {
	subject := models.Subject{}
	c.Bind(&subject)
	id, _ := strconv.Atoi(c.Param("id"))
	subject.Id = id

	err := repositories.UpdateSubject(database.DbConnection, subject)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": http.StatusText(http.StatusInternalServerError),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success update subject by id",
		"data":    subject,
	})
}

func DeleteSubjectById(c echo.Context) error {
	subject := models.Subject{}
	id, _ := strconv.Atoi(c.Param("id"))
	subject.Id = id

	err := repositories.DeleteSubject(database.DbConnection, subject)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": http.StatusText(http.StatusInternalServerError),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success delete subject by id",
		"data":    nil,
	})
}

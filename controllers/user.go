package controllers

import (
	"net/http"
	"strconv"

	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/database"
	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/models"
	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/repositories"
	"github.com/labstack/echo/v4"
)

func GetAllUser(c echo.Context) error {
	users, err := repositories.GetAllUser(database.DbConnection)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": http.StatusText(http.StatusInternalServerError),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all users",
		"data":    users,
	})
}

func CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	data, err := repositories.InsertUser(database.DbConnection, user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": http.StatusText(http.StatusInternalServerError),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create new user",
		"data":    data,
	})
}

func UpdateUserById(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	id, _ := strconv.Atoi(c.Param("id"))
	user.Id = id

	err := repositories.UpdateUser(database.DbConnection, user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": http.StatusText(http.StatusInternalServerError),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success update user by id",
		"data":    user,
	})
}

func DeleteUserById(c echo.Context) error {
	user := models.User{}
	id, _ := strconv.Atoi(c.Param("id"))
	user.Id = id

	err := repositories.DeleteUser(database.DbConnection, user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": http.StatusText(http.StatusInternalServerError),
			"data":    nil,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success delete user by id",
		"data":    nil,
	})
}

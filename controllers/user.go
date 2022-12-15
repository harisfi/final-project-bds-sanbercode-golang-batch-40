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
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get all users",
		"data":    users,
	})
}

func CreateUser(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := c.Validate(user); err != nil {
		return err
	}

	data, err := repositories.InsertUser(database.DbConnection, user)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create new user",
		"data":    data,
	})
}

func UpdateUserById(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := c.Validate(user); err != nil {
		return err
	}

	id, _ := strconv.Atoi(c.Param("id"))
	user.Id = id

	err := repositories.UpdateUser(database.DbConnection, user)

	if err != nil {
		return err
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
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success delete user by id",
		"data":    nil,
	})
}

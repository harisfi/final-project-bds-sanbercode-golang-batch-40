package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/database"
	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/middlewares"
	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/models"
	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/repositories"
	"github.com/harisfi/final-project-bds-sanbercode-golang-batch-40/utils"
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

func RegisterUser(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := c.Validate(user); err != nil {
		return err
	}

	pass, err := utils.HashPassword(user.Password)

	if err != nil {
		return err
	}

	user.Password = pass
	data, err := repositories.InsertUser(database.DbConnection, user)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success register new user",
		"data":    data,
	})
}

func LoginUser(c echo.Context) error {
	user := models.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := c.Validate(user); err != nil {
		return err
	}

	data, err := repositories.LoginUser(database.DbConnection, user)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success login user",
		"data":    data,
	})
}

func UpdateUserById(c echo.Context) error {
	user := models.User{}
	id, _ := strconv.Atoi(c.Param("id"))

	if middlewares.ExtractTokenUserId(c) != id {
		return errors.New("unauthorized")
	}

	if err := c.Bind(&user); err != nil {
		return err
	}
	if err := c.Validate(user); err != nil {
		return err
	}

	user.Id = id

	if user.Password != "" {
		pass, err := utils.HashPassword(user.Password)

		if err != nil {
			return err
		}

		user.Password = pass
	}

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

	if middlewares.ExtractTokenUserId(c) != id {
		return errors.New("unauthorized")
	}

	err := repositories.DeleteUser(database.DbConnection, user)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success delete user by id",
		"data":    nil,
	})
}

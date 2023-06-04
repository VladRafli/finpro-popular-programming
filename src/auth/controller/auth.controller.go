package controller

import (
	"my_kelurahan/auth/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	data, err := service.Login(c)

	if err != nil {
		return c.JSON(err.Code, echo.Map{
			"statusCode": err.Code,
			"message":    err.Message,
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"statusCode": http.StatusCreated,
		"message":    "Successfully created new user.",
		"data":       data,
	})
}

func Register(c echo.Context) error {
	data, err := service.Register(c)

	if err != nil {
		return c.JSON(err.Code, echo.Map{
			"statusCode": err.Code,
			"message":    err.Message,
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"statusCode": http.StatusCreated,
		"message":    "Successfully created new user.",
		"data":       data,
	})
}

func ForgotPassword(c echo.Context) error {
	data, err := service.ForgotPassword(c)

	if err != nil {
		return c.JSON(err.Code, echo.Map{
			"statusCode": err.Code,
			"message":    err.Message,
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"statusCode": http.StatusCreated,
		"message":    "Successfully created new user.",
		"data":       data,
	})
}

func ChangePassword(c echo.Context) error {
	data, err := service.ChangePassword(c)

	if err != nil {
		return c.JSON(err.Code, echo.Map{
			"statusCode": err.Code,
			"message":    err.Message,
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"statusCode": http.StatusCreated,
		"message":    "Successfully created new user.",
		"data":       data,
	})
}

func Logout(c echo.Context) error {
	data, err := service.Logout(c)

	if err != nil {
		return c.JSON(err.Code, echo.Map{
			"statusCode": err.Code,
			"message":    err.Message,
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"statusCode": http.StatusCreated,
		"message":    "Successfully created new user.",
		"data":       data,
	})
}

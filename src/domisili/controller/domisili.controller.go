package controller

import (
	"my_kelurahan/domisili/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	data, err := service.Create(c)

	if err != nil {
		return c.JSON(err.Code, echo.Map{
			"statusCode": err.Code,
			"message":    err.Message,
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"statusCode": http.StatusCreated,
		"message":    "Successfully created new Domisili.",
		"data":       data,
	})
}

func ReadAll(c echo.Context) error {
	data, err := service.ReadAll(c)

	if err != nil {
		return c.JSON(err.Code, echo.Map{
			"statusCode": err.Code,
			"message":    err.Message,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"statusCode": http.StatusOK,
		"message":    "Successfully retrieved all Domisili.",
		"data":       data,
	})
}

func Read(c echo.Context) error {
	data, err := service.Read(c)

	if err != nil {
		return c.JSON(err.Code, echo.Map{
			"statusCode": err.Code,
			"message":    err.Message,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"statusCode": http.StatusOK,
		"message":    "Successfully retrieved Domisili.",
		"data":       data,
	})
}

func Update(c echo.Context) error {
	data, err := service.Update(c)

	if err != nil {
		return c.JSON(err.Code, echo.Map{
			"statusCode": err.Code,
			"message":    err.Message,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"statusCode": http.StatusOK,
		"message":    "Successfully updated Domisili.",
		"data":       data,
	})
}

func Delete(c echo.Context) error {
	data, err := service.Delete(c)

	if err != nil {
		return c.JSON(err.Code, echo.Map{
			"statusCode": err.Code,
			"message":    err.Message,
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"statusCode": http.StatusOK,
		"message":    "Successfully deleted Domisili.",
		"data":       data,
	})
}
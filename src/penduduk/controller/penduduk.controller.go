package controller

import (
	"my_kelurahan/penduduk/service"
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
		"message":    "Successfully created new Penduduk.",
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
		"message":    "Successfully retrieved all Penduduk.",
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
		"message":    "Successfully retrieved Penduduk.",
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
		"message":    "Successfully updated Penduduk.",
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
		"message":    "Successfully deleted Penduduk.",
		"data":       data,
	})
}
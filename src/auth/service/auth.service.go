package service

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) (interface{}, *echo.HTTPError) {
	return nil, echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

func Register(c echo.Context) (interface{}, *echo.HTTPError) {
	return nil, echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

func ForgotPassword(c echo.Context) (interface{}, *echo.HTTPError) {
	return nil, echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

func ChangePassword(c echo.Context) (interface{}, *echo.HTTPError) {
	return nil, echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

func Logout(c echo.Context) (interface{}, *echo.HTTPError) {
	return nil, echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

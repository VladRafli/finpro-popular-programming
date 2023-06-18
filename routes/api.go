package routes

import (
	authControllers "my_kelurahan/auth/controller"
	"my_kelurahan/src/middlewares"

	"github.com/labstack/echo/v4"
)

func Init(app *echo.Echo) {
	apiGroup := app.Group("/api")
	authGroup := apiGroup.Group("/auth")
	authGroup.POST("/login", authControllers.Login)
	authGroup.POST("/register", authControllers.Register)
	authGroup.POST("/logout", authControllers.Logout, middlewares.ValidateJWT())
}
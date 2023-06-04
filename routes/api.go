package routes

import (
	authControllers "my_kelurahan/auth/controller"

	"github.com/labstack/echo/v4"
)

func Init(app *echo.Echo) {
	apiGroup := app.Group("/api")
	authGroup := apiGroup.Group("/auth")
	authGroup.POST("/login", authControllers.Login)
	authGroup.POST("/register", authControllers.Register)
	authGroup.POST("/forgot-password", authControllers.ForgotPassword)
	authGroup.POST("/change-password", authControllers.ChangePassword)
	authGroup.POST("/logout", authControllers.Logout)
}
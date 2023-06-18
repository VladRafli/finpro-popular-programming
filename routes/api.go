package routes

import (
	authControllers "my_kelurahan/auth/controller"
	rtControllers "my_kelurahan/rt/controller"
	"my_kelurahan/src/middlewares"

	"github.com/labstack/echo/v4"
)

func Init(app *echo.Echo) {
	apiGroup := app.Group("/api")

	authGroup := apiGroup.Group("/auth")
	authGroup.POST("/login", authControllers.Login)
	authGroup.POST("/register", authControllers.Register)
	authGroup.POST("/logout", authControllers.Logout, middlewares.ValidateJWT())

	rtGroup := apiGroup.Group("/rt")
	rtGroup.GET("", rtControllers.ReadAll, middlewares.ValidateJWT())
	rtGroup.GET("/:id", rtControllers.Read, middlewares.ValidateJWT())
	rtGroup.POST("", rtControllers.Create, middlewares.ValidateJWT())
	rtGroup.PUT("/:id", rtControllers.Update, middlewares.ValidateJWT())
	rtGroup.DELETE("/:id", rtControllers.Delete, middlewares.ValidateJWT())

	rwGroup := apiGroup.Group("/rw")
	rwGroup.GET("", rtControllers.ReadAll, middlewares.ValidateJWT())
	rwGroup.GET("/:id", rtControllers.Read, middlewares.ValidateJWT())
	rwGroup.POST("", rtControllers.Create, middlewares.ValidateJWT())
	rwGroup.PUT("/:id", rtControllers.Update, middlewares.ValidateJWT())
	rwGroup.DELETE("/:id", rtControllers.Delete, middlewares.ValidateJWT())

	domisiliGroup := apiGroup.Group("/domisili")
	domisiliGroup.GET("", rtControllers.ReadAll, middlewares.ValidateJWT())
	domisiliGroup.GET("/:id", rtControllers.Read, middlewares.ValidateJWT())
	domisiliGroup.POST("", rtControllers.Create, middlewares.ValidateJWT())
	domisiliGroup.PUT("/:id", rtControllers.Update, middlewares.ValidateJWT())
	domisiliGroup.DELETE("/:id", rtControllers.Delete, middlewares.ValidateJWT())

	pendudukGroup := apiGroup.Group("/penduduk")
	pendudukGroup.GET("", rtControllers.ReadAll, middlewares.ValidateJWT())
	pendudukGroup.GET("/:id", rtControllers.Read, middlewares.ValidateJWT())
	pendudukGroup.POST("", rtControllers.Create, middlewares.ValidateJWT())
	pendudukGroup.PUT("/:id", rtControllers.Update, middlewares.ValidateJWT())
	pendudukGroup.DELETE("/:id", rtControllers.Delete, middlewares.ValidateJWT())
}
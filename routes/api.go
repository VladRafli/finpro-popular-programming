package routes

import (
	authController "my_kelurahan/auth/controller"
	// rtControllers "my_kelurahan/rt/controller"
	// rwControllers "my_kelurahan/rw/controller"
	// domisiliControllers "my_kelurahan/domisili/controller"
	// pendudukControllers "my_kelurahan/penduduk/controller"
	"my_kelurahan/middlewares"

	"github.com/labstack/echo/v4"
)

func Init(app *echo.Echo) {
	apiGroup := app.Group("/api")

	authGroup := apiGroup.Group("/auth")
	authGroup.POST("login", authController.Login)
	authGroup.POST("register", authController.Register)
	authGroup.POST("/logout", authController.Logout, middlewares.ValidateJWT())

	// rtGroup := apiGroup.Group("/rt")
	// rtGroup.GET("", rtControllers.ReadAll, middlewares.ValidateJWT())
	// rtGroup.GET("/:id", rtControllers.Read, middlewares.ValidateJWT())
	// rtGroup.POST("", rtControllers.Create, middlewares.ValidateJWT())
	// rtGroup.PUT("/:id", rtControllers.Update, middlewares.ValidateJWT())
	// rtGroup.DELETE("/:id", rtControllers.Delete, middlewares.ValidateJWT())

	// rwGroup := apiGroup.Group("/rw")
	// rwGroup.GET("", rwControllers.ReadAll, middlewares.ValidateJWT())
	// rwGroup.GET("/:id", rwControllers.Read, middlewares.ValidateJWT())
	// rwGroup.POST("", rwControllers.Create, middlewares.ValidateJWT())
	// rwGroup.PUT("/:id", rwControllers.Update, middlewares.ValidateJWT())
	// rwGroup.DELETE("/:id", rwControllers.Delete, middlewares.ValidateJWT())

	// domisiliGroup := apiGroup.Group("/domisili")
	// domisiliGroup.GET("", domisiliControllers.ReadAll, middlewares.ValidateJWT())
	// domisiliGroup.GET("/:id", domisiliControllers.Read, middlewares.ValidateJWT())
	// domisiliGroup.POST("", domisiliControllers.Create, middlewares.ValidateJWT())
	// domisiliGroup.PUT("/:id", domisiliControllers.Update, middlewares.ValidateJWT())
	// domisiliGroup.DELETE("/:id", domisiliControllers.Delete, middlewares.ValidateJWT())

	// pendudukGroup := apiGroup.Group("/penduduk")
	// pendudukGroup.GET("", pendudukControllers.ReadAll, middlewares.ValidateJWT())
	// pendudukGroup.GET("/:id", pendudukControllers.Read, middlewares.ValidateJWT())
	// pendudukGroup.POST("", pendudukControllers.Create, middlewares.ValidateJWT())
	// pendudukGroup.PUT("/:id", pendudukControllers.Update, middlewares.ValidateJWT())
	// pendudukGroup.DELETE("/:id", pendudukControllers.Delete, middlewares.ValidateJWT())
}
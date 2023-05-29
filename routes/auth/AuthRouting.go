package auth

import (
	"github.com/gin-gonic/gin"
	"finpro/controllers"
)

func InitAuthRoute(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", controllers.PostLoginController)
		auth.POST("/register", controllers.PostRegisterController)
		auth.POST("/logout", controllers.PostLogoutController)
	}
}
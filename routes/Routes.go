package routes

import (
	"github.com/gin-gonic/gin"
	"finpro/routes/auth"
)

func Init(router *gin.Engine) {
	// Initialize the routes
	initRoutes(router)
}

func initRoutes(router *gin.Engine) {
	auth.InitAuthRoute(router)
}
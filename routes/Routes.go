package routes

import "github.com/gin-gonic/gin"

func Init(router *gin.Engine) {
	// Initialize the routes
	initRoutes(router)
}

func initRoutes(router *gin.Engine) {
	InitAuthRoute(router)
}
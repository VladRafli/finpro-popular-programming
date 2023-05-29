package main

import (
	"finpro/routes"
	"github.com/gin-gonic/gin"
)

func main() {
    app := gin.New()

    app.Use(gin.Logger())
    app.Use(gin.Recovery())
    routes.Init(app)

    app.Run(":5000")
}
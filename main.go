package main

import (
	"github.com/gin-gonic/gin"
    "finpro/routes"
)

func main() {
    app := gin.New()

    app.Use(gin.Logger())
    app.Use(gin.Recovery())
    routes.Init(app)

    app.Run(":5000")
}
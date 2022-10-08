package main

import (
	routers "eCommerce/route"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/gin-gonic/gin"
)

import "eCommerce/config"

func main() {
	app := gin.New()

	appConfig := config.AppConfig{}
	_ = env.Parse(&appConfig)

	// Setup route
	router := app.Group("/api")
	routers.SetupRouter(router)
	// Set domain
	domain := fmt.Sprintf(":%s", appConfig.AppPort)
	// run app
	_ = app.Run(domain)
}

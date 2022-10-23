package route

import (
	"eCommerce/controller"
	"github.com/gin-gonic/gin"
)

func HealthCheckRoutes(rootRoute *gin.RouterGroup) {
	healthCheckRouter := rootRoute.Group("/health-check")
	{
		healthCheckRouter.GET("/status", controller.HealthCheck)
	}
}

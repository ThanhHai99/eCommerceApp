package route

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(rootRoute *gin.RouterGroup) {
	HealthCheckRoutes(rootRoute)
	AuthenticateRoutes(rootRoute)
	ItemRoutes(rootRoute)
}

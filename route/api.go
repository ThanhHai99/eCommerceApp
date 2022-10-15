package route

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(rootRoute *gin.RouterGroup) {
	healthCheckRoutes(rootRoute)
	authenticateRoutes(rootRoute)
}

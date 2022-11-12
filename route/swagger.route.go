package route

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SwaggerRoutes(rootRoute *gin.RouterGroup) {
	swaggerRouter := rootRoute.Group("/swagger/docs")
	{
		swaggerRouter.GET("/", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}
}

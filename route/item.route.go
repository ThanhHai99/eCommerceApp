package route

import (
	"eCommerce/controller"
	"eCommerce/middleware"
	"github.com/gin-gonic/gin"
)

func ItemRoutes(rootRoute *gin.RouterGroup) {
	itemRouter := rootRoute.Group("/item", middleware.ValidateAccessToken())
	{
		itemRouter.GET("/", controller.GetAllItem)
		itemRouter.GET("/:id", controller.GetOneItem)
		itemRouter.POST("/", controller.CreateOneItem)
	}
}

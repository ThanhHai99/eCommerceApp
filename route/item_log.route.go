package route

import (
	"eCommerce/controller"
	"eCommerce/middleware"
	"github.com/gin-gonic/gin"
)

func ItemLogRoutes(rootRoute *gin.RouterGroup) {
	itemLogRouter := rootRoute.Group("/item-log", middleware.ValidateAccessToken())
	{
		itemLogRouter.GET("/", controller.GetAllItemLog)
		itemLogRouter.GET("/:id", controller.GetOneItemLog)
		itemLogRouter.PATCH("/:id", controller.UpdateItemLog)
		itemLogRouter.POST("/", controller.CreateOneItemLog)
		itemLogRouter.DELETE("/:id", controller.DeleteItemLog)
	}
}

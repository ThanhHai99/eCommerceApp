package route

import (
	"eCommerce/controller"
	"eCommerce/middleware"
	"eCommerce/validation"
	"github.com/gin-gonic/gin"
)

func ItemLogRoutes(rootRoute *gin.RouterGroup) {
	itemLogRouter := rootRoute.Group("/item-log", middleware.ValidateAccessToken())
	{
		itemLogRouter.GET("/", controller.GetAllItemLog)
		itemLogRouter.GET("/:id", controller.GetOneItemLog)
		itemLogRouter.PATCH("/:id", controller.UpdateItemLog)
		itemLogRouter.POST("/", validation.CreateItemLogBodyValidate(), controller.CreateOneItemLog)
		itemLogRouter.DELETE("/:id", controller.DeleteItemLog)
	}
}

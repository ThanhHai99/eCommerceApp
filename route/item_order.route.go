package route

import (
	"eCommerce/controller"
	"eCommerce/middleware"
	"eCommerce/validation"
	"github.com/gin-gonic/gin"
)

func ItemOrderRoutes(rootRoute *gin.RouterGroup) {
	itemOrderRouter := rootRoute.Group("/item-order", middleware.ValidateAccessToken())
	{
		itemOrderRouter.GET("/", controller.GetAllItemOrder)
		itemOrderRouter.GET("/:id", controller.GetOneItemOrder)
		itemOrderRouter.PATCH("/:id", controller.UpdateItemOrder)
		itemOrderRouter.POST("/", validation.CreateItemOrderBodyValidate(), controller.CreateOneItemOrder)
		itemOrderRouter.DELETE("/:id", controller.DeleteItemOrder)
	}
}

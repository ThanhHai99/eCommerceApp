package route

import (
	"eCommerce/controller"
	"eCommerce/middleware"
	"eCommerce/validation"
	"github.com/gin-gonic/gin"
)

func ItemRoutes(rootRoute *gin.RouterGroup) {
	itemRouter := rootRoute.Group("/item", middleware.ValidateAccessToken())
	{
		itemRouter.GET("/", controller.GetAllItem)
		itemRouter.GET("/:id", controller.GetOneItem)
		itemRouter.PATCH("/:id", controller.UpdateItem)
		itemRouter.POST("/", validation.CreateItemBodyValidate(), controller.CreateOneItem)
		itemRouter.DELETE("/:id", controller.DeleteItem)
	}
}

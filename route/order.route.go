package route

import (
	"eCommerce/controller"
	"eCommerce/middleware"
	"github.com/gin-gonic/gin"
)

func OrderRoutes(rootRoute *gin.RouterGroup) {
	orderRouter := rootRoute.Group("/order", middleware.ValidateAccessToken())
	{
		orderRouter.GET("/", controller.GetAllOrder)
		orderRouter.GET("/:id", controller.GetOneOrder)
		orderRouter.PATCH("/:id", controller.UpdateOrder)
		orderRouter.POST("/", controller.CreateOneOrder)
		orderRouter.DELETE("/:id", controller.DeleteOrder)
	}
}

package route

import (
	"eCommerce/controller"
	"eCommerce/middleware"
	"eCommerce/validation"
	"github.com/gin-gonic/gin"
)

func WarehouseLogRoutes(rootRoute *gin.RouterGroup) {
	warehouseLogRouter := rootRoute.Group("/warehouse-log", middleware.ValidateAccessToken())
	{
		warehouseLogRouter.GET("/", controller.GetAllWarehouseLog)
		warehouseLogRouter.GET("/:id", controller.GetOneWarehouseLog)
		warehouseLogRouter.PATCH("/:id", controller.UpdateWarehouseLog)
		warehouseLogRouter.POST("/", validation.CreateWarehouseLogBodyValidate(), controller.CreateOneWarehouseLog)
		warehouseLogRouter.DELETE("/:id", controller.DeleteWarehouseLog)
	}
}

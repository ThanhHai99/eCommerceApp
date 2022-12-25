package route

import (
	"eCommerce/controller"
	"eCommerce/middleware"
	"eCommerce/validation"
	"github.com/gin-gonic/gin"
)

func WarehouseRoutes(rootRoute *gin.RouterGroup) {
	warehouseRouter := rootRoute.Group("/warehouse", middleware.ValidateAccessToken())
	{
		warehouseRouter.GET("/", controller.GetAllWarehouse)
		warehouseRouter.GET("/:id", controller.GetOneWarehouse)
		warehouseRouter.PATCH("/:id", controller.UpdateWarehouse)
		warehouseRouter.POST("/", validation.CreateWarehouseBodyValidate(), controller.CreateOneWarehouse)
		warehouseRouter.DELETE("/:id", controller.DeleteWarehouse)
	}
}

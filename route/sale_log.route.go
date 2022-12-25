package route

import (
	"eCommerce/controller"
	"eCommerce/middleware"
	"eCommerce/validation"
	"github.com/gin-gonic/gin"
)

func SaleLogRoutes(rootRoute *gin.RouterGroup) {
	saleLogRouter := rootRoute.Group("/sale-log", middleware.ValidateAccessToken())
	{
		saleLogRouter.GET("/", controller.GetAllSaleLog)
		saleLogRouter.GET("/:id", controller.GetOneSaleLog)
		saleLogRouter.PATCH("/:id", controller.UpdateSaleLog)
		saleLogRouter.POST("/", validation.CreateSaleLogBodyValidate(), controller.CreateOneSaleLog)
		saleLogRouter.DELETE("/:id", controller.DeleteSaleLog)
	}
}

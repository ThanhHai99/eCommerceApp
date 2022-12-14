package route

import (
	"eCommerce/controller"
	"eCommerce/middleware"
	"eCommerce/validation"
	"github.com/gin-gonic/gin"
)

func SaleItemRoutes(rootRoute *gin.RouterGroup) {
	saleItemRouter := rootRoute.Group("/sale-item", middleware.ValidateAccessToken())
	{
		saleItemRouter.GET("/", controller.GetAllSaleItem)
		saleItemRouter.GET("/:id", controller.GetOneSaleItem)
		saleItemRouter.PATCH("/:id", controller.UpdateSaleItem)
		saleItemRouter.POST("/", validation.CreateSaleItemBodyValidate(), controller.CreateOneSaleItem)
		saleItemRouter.DELETE("/:id", controller.DeleteSaleItem)
	}
}

package route

import (
	"eCommerce/controller"
	"eCommerce/middleware"
	"github.com/gin-gonic/gin"
)

func SaleRoutes(rootRoute *gin.RouterGroup) {
	saleRouter := rootRoute.Group("/sale", middleware.ValidateAccessToken())
	{
		saleRouter.GET("/", controller.GetAllSale)
		saleRouter.GET("/:id", controller.GetOneSale)
		saleRouter.PATCH("/:id", controller.UpdateSale)
		saleRouter.POST("/", controller.CreateOneSale)
		saleRouter.DELETE("/:id", controller.DeleteSale)
	}
}

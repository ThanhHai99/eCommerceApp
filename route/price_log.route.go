package route

import (
	"eCommerce/controller"
	"eCommerce/middleware"
	"eCommerce/validation"
	"github.com/gin-gonic/gin"
)

func PriceLogRoutes(rootRoute *gin.RouterGroup) {
	priceLogRouter := rootRoute.Group("/price-log", middleware.ValidateAccessToken())
	{
		priceLogRouter.GET("/", controller.GetAllPriceLog)
		priceLogRouter.GET("/:id", controller.GetOnePriceLog)
		priceLogRouter.PATCH("/:id", controller.UpdatePriceLog)
		priceLogRouter.POST("/", validation.CreatePriceLogBodyValidate(), controller.CreateOnePriceLog)
		priceLogRouter.DELETE("/:id", controller.DeletePriceLog)
	}
}

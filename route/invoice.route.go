package route

import (
	"eCommerce/controller"
	"eCommerce/middleware"
	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(rootRoute *gin.RouterGroup) {
	invoiceRouter := rootRoute.Group("/invoice", middleware.ValidateAccessToken())
	{
		invoiceRouter.GET("/", controller.GetAllInvoice)
		invoiceRouter.GET("/:id", controller.GetOneInvoice)
		invoiceRouter.PATCH("/:id", controller.UpdateInvoice)
		invoiceRouter.POST("/", controller.CreateOneInvoice)
		invoiceRouter.DELETE("/:id", controller.DeleteInvoice)
	}
}

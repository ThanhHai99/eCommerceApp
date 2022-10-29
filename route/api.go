package route

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(rootRoute *gin.RouterGroup) {
	HealthCheckRoutes(rootRoute)
	AuthenticateRoutes(rootRoute)
	ItemRoutes(rootRoute)
	InvoiceRoutes(rootRoute)
	OrderRoutes(rootRoute)
	SaleRoutes(rootRoute)
	UserRoutes(rootRoute)
	CategoryLogRoutes(rootRoute)
	ItemLogRoutes(rootRoute)
	ItemOrderRoutes(rootRoute)
	PriceLogRoutes(rootRoute)
	RoleRoutes(rootRoute)
	SaleItemRoutes(rootRoute)
	SaleLogRoutes(rootRoute)
	WarehouseRoutes(rootRoute)
	WarehouseLogRoutes(rootRoute)
}

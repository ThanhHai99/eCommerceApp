package route

import (
	"eCommerce/controller"
	"eCommerce/middleware"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(rootRoute *gin.RouterGroup) {
	categoryRouter := rootRoute.Group("/category", middleware.ValidateAccessToken())
	{
		categoryRouter.GET("/", controller.GetAllCategory)
		categoryRouter.GET("/:id", controller.GetOneCategory)
		categoryRouter.PATCH("/:id", controller.UpdateCategory)
		categoryRouter.POST("/", controller.CreateOneCategory)
		categoryRouter.DELETE("/:id", controller.DeleteCategory)
	}
}

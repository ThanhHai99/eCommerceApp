package route

import (
	"eCommerce/controller"
	"eCommerce/middleware"
	"github.com/gin-gonic/gin"
)

func CategoryLogRoutes(rootRoute *gin.RouterGroup) {
	categoryLogRouter := rootRoute.Group("/category-log", middleware.ValidateAccessToken())
	{
		categoryLogRouter.GET("/", controller.GetAllCategoryLog)
		categoryLogRouter.GET("/:id", controller.GetOneCategoryLog)
		categoryLogRouter.PATCH("/:id", controller.UpdateCategoryLog)
		categoryLogRouter.POST("/", controller.CreateOneCategoryLog)
		categoryLogRouter.DELETE("/:id", controller.DeleteCategoryLog)
	}
}

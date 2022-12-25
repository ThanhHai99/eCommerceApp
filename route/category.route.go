package route

import (
	"eCommerce/controller"
	"eCommerce/middleware"
	"eCommerce/validation"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(rootRoute *gin.RouterGroup) {
	categoryRouter := rootRoute.Group("/category", middleware.ValidateAccessToken())
	{
		categoryRouter.GET("/", controller.GetAllCategory)
		categoryRouter.GET("/:id", controller.GetOneCategory)
		categoryRouter.PATCH("/:id", validation.UpdateCategoryBodyValidate(), controller.UpdateCategory)
		categoryRouter.POST("/", validation.CreateCategoryBodyValidate(), controller.CreateOneCategory)
		categoryRouter.DELETE("/:id", controller.DeleteCategory)
	}
}
